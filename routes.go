package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.Handler
}

var routes = []Route{
	Route{ // This is a special route
		"AuthBoss",
		"",
		"/auth",
		nil,
	},
	Route{
		"Index",
		"GET",
		"/",
		authProtect(Index, true),
	},
	Route{
		"Issue",
		"GET",
		"/i/{id}",
		authProtect(Issue, false),
	},
}
