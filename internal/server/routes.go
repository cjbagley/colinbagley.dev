package server

import (
	"fmt"
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"net/http"
)

// Route holds information for a given route
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// GetRoutes returns all available routes for the website
func GetRoutes() []Route {
	routes := []Route{
		{http.MethodGet, "/", handleIndex},
		{http.MethodGet, "/articles", handleArticles},
	}

	articles := data.GetArticles()

	if len(articles) > 0 {
		for _, article := range articles {
			h := handleArticle(article)
			routes = append(routes, Route{http.MethodGet, "/" + article.URL, h})
		}
	}

	return routes
}

// AddRoutesToMux adds all available routes to a servemux instance
func AddRoutesToMux(mux *http.ServeMux) {
	routes := GetRoutes()
	for _, r := range routes {
		mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Path), r.Handler)
	}
}
