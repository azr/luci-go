// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package mutate

import (
	"testing"

	"github.com/luci/gae/filter/featureBreaker"
	"github.com/luci/gae/impl/memory"
	"github.com/luci/gae/service/datastore"
	"github.com/luci/luci-go/appengine/cmd/dm/model"
	"github.com/luci/luci-go/appengine/cmd/dm/types"
	"github.com/luci/luci-go/appengine/tumble"
	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
)

func TestAddBackDep(t *testing.T) {
	t.Parallel()

	Convey("AddBackDep", t, func() {
		c := memory.Use(context.Background())
		ds := datastore.Get(c)

		abd := &AddBackDep{
			Dep: &model.FwdEdge{
				From: types.NewAttemptID("quest|fffffffe"),
				To:   types.NewAttemptID("to|fffffffe"),
			},
		}

		Convey("Root", func() {
			So(abd.Root(c).String(), ShouldEqual, `dev~app::/BackDepGroup,"to|fffffffe"`)
		})

		Convey("RollForward", func() {
			bdg, bd := abd.Dep.Back(c)
			So(bd.Propagated, ShouldBeFalse)

			Convey("attempt finished", func() {
				bdg.AttemptFinished = true
				So(ds.Put(bdg), ShouldBeNil)

				Convey("no need completion", func() {
					muts, err := abd.RollForward(c)
					So(err, ShouldBeNil)
					So(muts, ShouldBeNil)

					So(ds.GetMulti([]interface{}{bdg, bd}), ShouldBeNil)
					So(bd.Edge(), ShouldResembleV, abd.Dep)
					So(bd.Propagated, ShouldBeTrue)
				})

				Convey("need completion", func() {
					abd.NeedsAck = true
					muts, err := abd.RollForward(c)
					So(err, ShouldBeNil)
					So(muts, ShouldResembleV, []tumble.Mutation{
						&AckFwdDep{abd.Dep, true}})

					So(ds.GetMulti([]interface{}{bdg, bd}), ShouldBeNil)
					So(bd.Edge(), ShouldResembleV, abd.Dep)
					So(bd.Propagated, ShouldBeTrue)
				})
			})

			Convey("attempt not finished, need completion", func() {
				ex, err := ds.Exists(ds.KeyForObj(bdg))
				So(err, ShouldBeNil)
				So(ex, ShouldBeFalse)

				abd.NeedsAck = true
				muts, err := abd.RollForward(c)
				So(err, ShouldBeNil)
				So(muts, ShouldResembleV, []tumble.Mutation{
					&AckFwdDep{abd.Dep, false}})

				// Note that bdg was created as a side effect.
				So(ds.GetMulti([]interface{}{bdg, bd}), ShouldBeNil)
				So(bd.Edge(), ShouldResembleV, abd.Dep)
				So(bd.Propagated, ShouldBeFalse)
				So(bdg.AttemptFinished, ShouldBeFalse)
			})

			Convey("failure", func() {
				c, fb := featureBreaker.FilterRDS(c, nil)
				fb.BreakFeatures(nil, "PutMulti")

				_, err := abd.RollForward(c)
				So(err, ShouldErrLike, `feature "PutMulti" is broken`)
			})
		})
	})
}