package server

import (
	"fmt"
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func GetRoutes() []Route {
	routes := []Route{
		{http.MethodGet, "/", HandleIndex},
		{http.MethodGet, "/articles", HandleArticles},
	}

	articles := data.GetArticles()

	if len(articles) > 0 {
		for _, article := range articles {
			h := HandleArticle(article)
			routes = append(routes, Route{http.MethodGet, "/" + article.URL, h})
		}
	}

	return routes
}

func AddRoutesToMux(mux *http.ServeMux) {
	routes := GetRoutes()
	for _, r := range routes {
		mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Path), r.Handler)
	}
}
