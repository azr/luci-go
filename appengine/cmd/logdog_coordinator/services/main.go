// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package module

import (
	"net/http"

	"github.com/luci/luci-go/appengine/gaemiddleware"
	"github.com/luci/luci-go/appengine/logdog/coordinator"
	"github.com/luci/luci-go/appengine/logdog/coordinator/config"
	"github.com/luci/luci-go/appengine/logdog/coordinator/endpoints/registration"
	"github.com/luci/luci-go/appengine/logdog/coordinator/endpoints/services"
	registrationPb "github.com/luci/luci-go/common/api/logdog_coordinator/registration/v1"
	servicesPb "github.com/luci/luci-go/common/api/logdog_coordinator/services/v1"
	"github.com/luci/luci-go/server/prpc"
	"github.com/luci/luci-go/server/router"

	// Include mutations package so its Mutations will register with tumble via
	// init().
	_ "github.com/luci/luci-go/appengine/logdog/coordinator/mutations"
)

// base returns the root middleware chain.
func base() router.MiddlewareChain {
	return gaemiddleware.BaseProd().Extend(
		coordinator.WithProdServices,
		config.WithConfig,
	)
}

// Run installs and executes this site.
func init() {
	r := router.New()

	// Setup Cloud Endpoints.
	svr := prpc.Server{}
	servicesPb.RegisterServicesServer(&svr, services.New())
	registrationPb.RegisterRegistrationServer(&svr, registration.New())

	// Standard HTTP endpoints.
	svr.InstallHandlers(r, base())

	http.Handle("/", r)
}
