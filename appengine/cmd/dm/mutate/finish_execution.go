// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package mutate

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/luci/gae/filter/txnBuf"
	"github.com/luci/gae/service/datastore"

	"github.com/luci/luci-go/appengine/cmd/dm/distributor"
	"github.com/luci/luci-go/appengine/cmd/dm/model"
	"github.com/luci/luci-go/appengine/tumble"
	dm "github.com/luci/luci-go/common/api/dm/service/v1"
)

// FinishExecution records the final state of the Execution, and advances the
// Attempt state machine.
type FinishExecution struct {
	EID    *dm.Execution_ID
	Result *distributor.TaskResult
}

// Root implements tumble.Mutation
func (f *FinishExecution) Root(c context.Context) *datastore.Key {
	return model.ExecutionKeyFromID(c, f.EID)
}

// shouldRetry loads the quest for this attempt, to determine if the attempt can
// be retried. As a side-effect, it increments the RetryState counter for the
// indicated failure type.
//
// If stat is not a retryable AbnormalFinish_Status, this will panic.
func shouldRetry(c context.Context, a *model.Attempt, stat dm.AbnormalFinish_Status) (retry bool, err error) {
	if !stat.CouldRetry() {
		return
	}
	q := model.QuestFromID(a.ID.Quest)
	dsNoTxn := txnBuf.GetNoTxn(c)
	if err = dsNoTxn.Get(q); err != nil {
		return
	}
	var cur, max uint32
	switch stat {
	case dm.AbnormalFinish_FAILED:
		cur, max = a.RetryState.Failed, q.Desc.Meta.Retry.Failed
		a.RetryState.Failed++
	case dm.AbnormalFinish_CRASHED:
		cur, max = a.RetryState.Crashed, q.Desc.Meta.Retry.Crashed
		a.RetryState.Crashed++
	case dm.AbnormalFinish_EXPIRED:
		cur, max = a.RetryState.Expired, q.Desc.Meta.Retry.Expired
		a.RetryState.Expired++
	case dm.AbnormalFinish_TIMED_OUT:
		cur, max = a.RetryState.TimedOut, q.Desc.Meta.Retry.TimedOut
		a.RetryState.TimedOut++
	default:
		panic(fmt.Errorf("do not know how to retry %q", stat))
	}
	retry = cur < max
	return
}

// RollForward implements tumble.Mutation
func (f *FinishExecution) RollForward(c context.Context) (muts []tumble.Mutation, err error) {
	a := model.AttemptFromID(f.EID.AttemptID())
	e := model.ExecutionFromID(c, f.EID)

	ds := datastore.Get(c)
	if err = ds.Get(a, e); err != nil {
		return
	}

	if a.State != dm.Attempt_EXECUTING || a.CurExecution != f.EID.Id || e.State.Terminal() {
		return
	}

	if f.Result.AbnormalFinish == nil && e.State != dm.Execution_STOPPING {
		f.Result.AbnormalFinish = &dm.AbnormalFinish{
			Status: dm.AbnormalFinish_FAILED,
			Reason: fmt.Sprintf("distributor finished execution while it was in the %s state.", e.State),
		}
	}

	if ab := f.Result.AbnormalFinish; ab != nil {
		if err = e.ModifyState(c, dm.Execution_ABNORMAL_FINISHED); err != nil {
			return
		}
		e.AbnormalFinish = *ab

		var retry bool
		if retry, err = shouldRetry(c, a, ab.Status); err != nil {
			return
		} else if retry {
			if err = a.ModifyState(c, dm.Attempt_SCHEDULING); err != nil {
				return
			}
			a.DepMap.Reset()
			muts = append(muts, &ScheduleExecution{&a.ID})
		} else {
			// ran out of retries, or non-retriable error type
			if err = a.ModifyState(c, dm.Attempt_ABNORMAL_FINISHED); err != nil {
				return
			}
			a.AbnormalFinish = *ab
		}
	} else {
		if err = e.ModifyState(c, dm.Execution_FINISHED); err != nil {
			return
		}
		e.ResultPersistentState = f.Result.PersistentState

		a.PersistentState = f.Result.PersistentState
		a.RetryState.Reset()

		if a.DepMap.Size() > 0 {
			if err = a.ModifyState(c, dm.Attempt_WAITING); err != nil {
				return
			}
		} else {
			if err = a.ModifyState(c, dm.Attempt_FINISHED); err != nil {
				return
			}
			muts = append(muts, &RecordCompletion{f.EID.AttemptID()})
		}
	}

	// best-effort reset execution timeout
	_ = ResetExecutionTimeout(c, e)

	err = ds.Put(a, e)
	return
}

// FinishExecutionFn is the implementation of distributor.FinishExecutionFn.
// It's defined here to avoid a circular dependency.
func FinishExecutionFn(c context.Context, eid *dm.Execution_ID, rslt *distributor.TaskResult) ([]tumble.Mutation, error) {
	return []tumble.Mutation{&FinishExecution{EID: eid, Result: rslt}}, nil
}

// NewFinishExecutionAbnormal is a shorthand to make a FinishExecution mutation
// with some abnomal result.
func NewFinishExecutionAbnormal(eid *dm.Execution_ID, status dm.AbnormalFinish_Status, reason string) *FinishExecution {
	return &FinishExecution{
		eid, &distributor.TaskResult{
			AbnormalFinish: &dm.AbnormalFinish{
				Status: status, Reason: reason}}}
}

func init() {
	tumble.Register((*FinishExecution)(nil))
}
