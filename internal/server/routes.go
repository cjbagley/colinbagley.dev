package server

import (
	"encoding/json"
	"fmt"
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"net/http"
	"os"
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

	var articles []data.Article
	if articleJson, err := os.ReadFile("internal/data/articles.json"); err == nil {
		json.Unmarshal(articleJson, &articles)
	}

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
