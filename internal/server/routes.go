package server

import (
	"fmt"
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"net/http"
)

type Route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func AddRoutes(mux *http.ServeMux) {
	routes := []Route{
		{http.MethodGet, "/articles", HandleArticles},
		{http.MethodGet, "/", HandleIndex},
	}

	articles := data.GetArticles()

	if len(articles) > 0 {
		for _, article := range articles {
			h := HandleArticle(article)
			routes = append(routes, Route{http.MethodGet, "/" + article.Url, h})
		}
	}

	for _, r := range routes {
		mux.HandleFunc(fmt.Sprintf("%s %s", r.method, r.path), r.handler)
	}
}
