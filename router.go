package main

import (
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"

	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/register"

	"github.com/rdnelson/reclus/backends"
	"github.com/rdnelson/reclus/config"
	"github.com/rdnelson/reclus/log"
)

var (
	authManager = authboss.New()
)

func setupAuth(db backends.Database) error {

	if err := InitializeCookieStore(); err != nil {
		return err
	}

	if err := InitializeSessionStore(); err != nil {
		return err
	}

	log.Log.Debug("Setting up authentication manager")

	authManager.Storer = NewUserRepo(db)
	authManager.SessionStoreMaker = NewSessionStore
	authManager.CookieStoreMaker = NewCookieStore
	authManager.MountPath = "/auth"
	authManager.RootURL = config.Cfg.Server.DisplayName
	authManager.LogWriter = log.LogWriter{log.Log}

	authManager.XSRFName = "csrf_token"
	authManager.XSRFMaker = func(_ http.ResponseWriter, req *http.Request) string {
		return nosurf.Token(req)
	}

	return authManager.Init()
}

func RouteLogger(handler http.Handler, name string) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		log.Log.Infof(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		handler := route.HandlerFunc

		handler = RouteLogger(handler, route.Name)

		// This needs an exception because authManager needs to be initialized
		// and therefore can't be put into the routing table at startup
		if route.Name == "AuthBoss" {
			router.
				PathPrefix(route.Pattern).
				Name(route.Name).
				Handler(context.ClearHandler(authManager.NewRouter()))

			continue
		}

		if route.Pattern != "" {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		} else {
			log.Log.Warningf(
				"No route set for '%s'",
				route.Name)
		}
	}

	return router
}
