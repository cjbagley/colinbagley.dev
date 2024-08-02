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
	path := "./internal/data/articles.json"
	if testing.Testing() {
		path = "../../internal/data/articles.json"
	}

	var articles []Article
	if articleJson, err := os.ReadFile(path); err == nil {
		err := json.Unmarshal(articleJson, &articles)
		if err != nil {
			// Still allow website to function
			return articles
		}
	}

	return articles
}
