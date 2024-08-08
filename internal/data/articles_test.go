package data

import (
	"encoding/json"
	"testing"
)

func TestArticles(t *testing.T) {
	rawJSON := []byte(`
		[
		  {
			"url": "article-1",
			"title": "Article 1",
			"template": "article-1"
		  },
		  {
			"url": "article-2",
			"title": "Article 2",
			"template": "article-2"
		  }
		]
	`)

	var articles []Article
	err := json.Unmarshal(rawJSON, &articles)
	if err != nil {
		t.Errorf("Could not unmarshal json: %v", err)
		return
	}

	if len(articles) != 2 {
		t.Errorf("Articles should have 2 items, got %v", len(articles))
	}
}
