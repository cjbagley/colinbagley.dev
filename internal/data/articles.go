// Package data houses structs and methods for website data
package data

import (
	"encoding/json"
	"os"
	"testing"
)

// Article data structure
type Article struct {
	URL       string `json:"url"`
	Title     string `json:"title"`
	Published string `json:"published"`
	Updated   string `json:"updated"`
	Template  string `json:"template"`
	Favourite bool   `json:"favourite"`
}

// GetArticles returns all available articles for the website
func GetArticles() []Article {
	path := "./internal/data/articles.json"
	if testing.Testing() {
		path = "../../internal/data/articles.json"
	}

	var articles []Article
	if articleJSON, err := os.ReadFile(path); err == nil {
		err := json.Unmarshal(articleJSON, &articles)
		if err != nil {
			// Still allow website to function
			return articles
		}
	}

	return articles
}
