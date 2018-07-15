package main

// Display a single article
func GetArticleHelper(articles []Article, id string) Article {
	for _, article := range articles {
		if article.Id == id {
			return article
		}
	}

	return Article{}
}
