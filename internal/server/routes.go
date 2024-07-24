package server

import (
	"fmt"
	"net/http"
)

type Route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func AddRoutes(mux *http.ServeMux) {
	routes := []Route{
		{http.MethodGet, "/", HandleIndex},
	}

	for _, r := range routes {
		mux.HandleFunc(fmt.Sprintf("%s %s", r.method, r.path), r.handler)
	}
}
