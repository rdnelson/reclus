package main

import (
	"net/http"

	"github.com/rdnelson/reclus/log"
)

type authHandler struct {
	handler http.HandlerFunc
	groups  []string
}

func authProtect(f http.HandlerFunc) authHandler {
	return authHandler{f, nil}
}

func authProtectGroups(f http.HandlerFunc, groups []string) authHandler {
	return authHandler{f, groups}
}

func (a authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := authManager.CurrentUser(w, r)
	if err != nil {
		log.Log.Warnf("Failed to fetch current user: %v", err)
	}

	if err != nil || u == nil {
		log.Log.Debugf("Redirecting unauthorized user to log.Login page. Return URL: '%s'", r.URL.String())

		cookies := NewCookieStore(w, r)
		cookies.Put("return_url", r.URL.String())

		http.Redirect(w, r, "/auth/log.Login", http.StatusFound)
	} else {
		a.handler(w, r)
	}
}
