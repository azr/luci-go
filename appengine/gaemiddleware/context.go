// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package gaemiddleware

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"

	"github.com/luci/gae/filter/dscache"
	"github.com/luci/gae/impl/prod"
	"github.com/luci/luci-go/appengine/gaeauth/client"
	"github.com/luci/luci-go/appengine/gaeauth/server"
	"github.com/luci/luci-go/appengine/gaeauth/server/gaesigner"
	"github.com/luci/luci-go/appengine/gaesecrets"
	"github.com/luci/luci-go/appengine/gaesettings"
	"github.com/luci/luci-go/appengine/tsmon"
	"github.com/luci/luci-go/common/cacheContext"
	"github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/server/auth"
	"github.com/luci/luci-go/server/middleware"
	"github.com/luci/luci-go/server/proccache"
	"github.com/luci/luci-go/server/router"
	"github.com/luci/luci-go/server/settings"
)

var (
	// globalProcessCache holds state cached between requests.
	globalProcessCache = &proccache.Cache{}

	// globalSettings holds global app settings lazily updated from the datastore.
	globalSettings = settings.New(gaesettings.Storage{})

	// globalAuthDBCache knows how to fetch auth.DB from datastore and keep it
	// in local memory cache. Used in prod contexts only.
	globalAuthDBCache = auth.NewDBCache(server.GetAuthDB)

	// globalTsMonState holds state related to time series monitoring.
	globalTsMonState = &tsmon.State{}
)

// WithProd installs the set of standard production AppEngine services:
//   * github.com/luci/luci-go/common/logging (set default level to debug).
//   * github.com/luci/gae/impl/prod (production appengine services)
//   * github.com/luci/luci-go/appengine/gaeauth/client (appengine urlfetch transport)
//   * github.com/luci/luci-go/server/proccache (in process memory cache)
//   * github.com/luci/luci-go/server/settings (global app settings)
//   * github.com/luci/luci-go/appengine/gaesecrets (access to secret keys in datastore)
//   * github.com/luci/luci-go/appengine/gaeauth/server/gaesigner (RSA signer)
//   * github.com/luci/luci-go/appengine/gaeauth/server/auth (user groups database)
func WithProd(c context.Context, req *http.Request) context.Context {
	// These are needed to use fetchCachedSettings.
	c = logging.SetLevel(c, logging.Debug)
	c = prod.Use(c, req)
	c = settings.Use(c, globalSettings)

	// Fetch and apply configuration stored in the datastore.
	settings := fetchCachedSettings(c)
	c = logging.SetLevel(c, settings.LoggingLevel)
	if !settings.DisableDSCache {
		c = dscache.AlwaysFilterRDS(c, nil)
	}

	// The rest of the service may use applied configuration.
	c = proccache.Use(c, globalProcessCache)
	c = client.UseAnonymousTransport(c)
	c = gaesecrets.Use(c, nil)
	c = gaesigner.Use(c)
	c = auth.UseDB(c, globalAuthDBCache)
	return cacheContext.Wrap(c)
}

// ProdServices is a middleware that installs the set of standard production
// AppEngine services by calling WithProd.
func ProdServices(c *router.Context, next router.Handler) {
	c.Context = WithProd(c.Context, c.Request)
	next(c)
}

// BaseProd returns a list of middleware: WithProd middleware, a panic catcher if this
// is not a devserver, and the monitoring middleware.
func BaseProd() router.MiddlewareChain {
	if appengine.IsDevAppServer() {
		return router.NewMiddlewareChain(ProdServices, globalTsMonState.Middleware)
	}
	return router.NewMiddlewareChain(
		ProdServices, middleware.WithPanicCatcher, globalTsMonState.Middleware,
	)
}
