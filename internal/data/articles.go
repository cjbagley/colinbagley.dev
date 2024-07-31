package data

import (
	"encoding/json"
	"os"
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
	if articleJson, err := os.ReadFile("internal/data/articles.json"); err == nil {
		json.Unmarshal(articleJson, &articles)
	}

	return articles
}
