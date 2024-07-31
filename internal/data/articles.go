package data

import (
	"encoding/json"
	"os"
	"testing"
)

type Article struct {
	Url       string `json:"url"`
	Title     string `json:"title"`
	Published string `json:"published"`
	Template  string `json:"template"`
	Favourite bool   `json:"favourite"`
}

func GetArticles() []Article {
	var articles []Article
	path := "./internal/data/articles.json"
	if testing.Testing() {
		path = "../../internal/data/articles.json"
	}

	if articleJson, err := os.ReadFile(path); err == nil {
		json.Unmarshal(articleJson, &articles)
	}

	return articles
}
