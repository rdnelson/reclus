package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	Prefix      string
	HandlerFunc http.Handler
}

var routes = []Route{
	Route{
		"Index",
		[]string{"GET"},
		"/",
		"",
		authProtect(Index, true),
	},
	Route{
		"AuthBoss",
		[]string{"GET", "POST"},
		"",
		"/auth",
		nil,
	},
}
