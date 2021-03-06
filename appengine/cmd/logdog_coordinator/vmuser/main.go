// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/luci/luci-go/appengine/gaemiddleware"
	"github.com/luci/luci-go/appengine/logdog/coordinator"
	"github.com/luci/luci-go/appengine/logdog/coordinator/config"
	"github.com/luci/luci-go/appengine/logdog/coordinator/endpoints/admin"
	"github.com/luci/luci-go/appengine/logdog/coordinator/endpoints/logs"
	"github.com/luci/luci-go/appengine/logdog/coordinator/endpoints/registration"
	"github.com/luci/luci-go/appengine/logdog/coordinator/endpoints/services"
	adminPb "github.com/luci/luci-go/common/api/logdog_coordinator/admin/v1"
	logsPb "github.com/luci/luci-go/common/api/logdog_coordinator/logs/v1"
	registrationPb "github.com/luci/luci-go/common/api/logdog_coordinator/registration/v1"
	servicesPb "github.com/luci/luci-go/common/api/logdog_coordinator/services/v1"
	log "github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/server/discovery"
	"github.com/luci/luci-go/server/prpc"
	"github.com/luci/luci-go/server/router"
	"golang.org/x/net/context"
	"google.golang.org/appengine"

	// Include mutations package so its Mutations will register with tumble via
	// init().
	_ "github.com/luci/luci-go/appengine/logdog/coordinator/mutations"
)

// base returns the root middleware chain.
func base(installConfig bool) router.MiddlewareChain {
	m := gaemiddleware.BaseProd().Extend(coordinator.WithProdServices)
	if installConfig {
		m = m.Extend(config.WithConfig)
	}
	return m
}

// Run installs and executes this site.
func main() {
	r := router.New()

	// Setup Cloud Endpoints.
	svr := prpc.Server{
		AccessControl: accessControl,
	}
	adminPb.RegisterAdminServer(&svr, admin.New())
	servicesPb.RegisterServicesServer(&svr, services.New())
	registrationPb.RegisterRegistrationServer(&svr, registration.New())
	logsPb.RegisterLogsServer(&svr, logs.New())
	discovery.Enable(&svr)

	// Standard HTTP endpoints.
	gaemiddleware.InstallHandlers(r, base(false))
	svr.InstallHandlers(r, base(true))

	// Redirect "/" to "/app/".
	r.GET("/", router.MiddlewareChain{}, func(c *router.Context) {
		http.Redirect(c.Writer, c.Request, "/app/", http.StatusFound)
	})

	http.Handle("/", r)
	appengine.Main()
}

func accessControl(c context.Context, origin string) bool {
	cfg, err := config.Load(c)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to get config for access control check.")
		return false
	}

	ccfg := cfg.GetCoordinator()
	if ccfg == nil {
		return false
	}

	for _, o := range ccfg.RpcAllowOrigins {
		if o == origin {
			return true
		}
	}
	return false
}
