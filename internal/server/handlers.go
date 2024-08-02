package server

import (
	"github.com/cjbagley/colinbagley.dev/internal/data"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	WriteHttpResponse(w, &contentPageTemplates{path: "index.gohtml", title: "Hello!"})
}

//goland:noinspection GoUnusedParameter
func HandleArticles(w http.ResponseWriter, r *http.Request) {
	articles := data.GetArticles()

	WriteHttpResponse(w, &contentPageTemplates{path: "articles.gohtml", title: "Articles", articles: articles})
}

func HandleArticle(article data.Article) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		WriteHttpResponse(w, &articlePageTemplates{article: article})
	}

	return fn
}
