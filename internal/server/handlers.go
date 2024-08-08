package server

import (
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		WriteHTTPResponse(w, &contentPageTemplates{path: "404.gohtml", title: "Page Not Found"})
		return
	}

	WriteHTTPResponse(w, &contentPageTemplates{path: "index.gohtml", title: "Hello!"})
}

func HandleArticles(w http.ResponseWriter, _ *http.Request) {
	articles := data.GetArticles()

	WriteHTTPResponse(w, &contentPageTemplates{path: "articles.gohtml", title: "Articles", articles: articles})
}

func HandleArticle(article data.Article) http.HandlerFunc {
	fn := func(w http.ResponseWriter, _ *http.Request) {
		WriteHTTPResponse(w, &articlePageTemplates{article: article})
	}

	return fn
}
