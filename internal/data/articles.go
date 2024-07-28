package data

type Article struct {
	Url       string `json:"url"`
	Title     string `json:"title"`
	Template  string `json:"template"`
	Favourite bool   `json:"favourite"`
}

func GetArticleMap(articles []Article) map[string]Article {
	articleMap := make(map[string]Article)
	for _, article := range articles {
		if article.Url == "" {
			continue
		}
		articleMap[article.Url] = article
	}

	return articleMap
}
