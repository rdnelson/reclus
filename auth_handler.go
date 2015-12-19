package main

import (
	"net/http"

	"github.com/rdnelson/reclus/datamodel"
	"github.com/rdnelson/reclus/log"
)

type AuthHttpHandler func(http.ResponseWriter, *http.Request, *datamodel.User)

type authHandler struct {
	handler AuthHttpHandler
	full    bool
	groups  []string
}

func authProtect(f AuthHttpHandler, full bool) authHandler {
	return authHandler{f, full, nil}
}

func authProtectGroups(f AuthHttpHandler, full bool, groups []string) authHandler {
	return authHandler{f, full, groups}
}

func (a authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := authManager.CurrentUser(w, r)

	if err != nil {
		log.Log.Warnf("Failed to fetch current user: %v", err)
	}

	if err != nil || u == nil {
		log.Log.Debugf("Redirecting unauthorized user to login page. Return URL: '%s'", r.URL.String())

		cookies := NewCookieStore(w, r)
		cookies.Put("return_url", r.URL.String())

		http.Redirect(w, r, "/auth/login", http.StatusFound)
	} else {

		usr := u.(*datamodel.User)
		if a.full {
			oldUsr := usr

			switch authManager.Storer.(type) {
			case AuthUserRepo:
				usr, err = authManager.Storer.(AuthUserRepo).GetUser(usr)
				break
			}

			if err != nil {
				usr = oldUsr
			}
		}

		a.handler(w, r, usr)
	}
}
