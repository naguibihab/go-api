package main

import (
	"strings"
)

// Get a single article from array by id
func GetArticleHelper(articles []Article, id string) Article {
	for _, article := range articles {
		if article.Id == id {
			return article
		}
	}

	return Article{}
}

// Get all articles that contain this date and tag
func GetArticlesByDateAndTagHelper(articles []Article, searchForDate string, searchForTag string) TagSearchResponse {
	result := TagSearchResponse{
		Tag:         searchForTag,
		Count:       0,
		Articles:    make([]string, 0),
		RelatedTags: make([]*Tag, 0),
	}

	articlesOnDate := getArticlesByDate(articles, searchForDate)
	for _, article := range articlesOnDate {
		for tagIndex, tag := range article.Tags {
			if tag.Name == searchForTag {
				// Increment the count
				result.Count++
				// Add id of article to our result
				result.Articles = append(result.Articles, article.Id)
				// Remove the tag we're looking for and add the other tags to related tags
				result.RelatedTags = append(article.Tags[:tagIndex], article.Tags[tagIndex+1:]...)
			}
		}
	}

	return result
}

func DeleteArticleHelper(articles []Article, findId string) []Article {
	for index, article := range articles {
		if article.Id == findId {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
		return articles
	}
	return articles
}

/// ---------- private functions ---------------- //

// Get all articles for a particular date
func getArticlesByDate(articles []Article, date string) []Article {
	var results []Article
	for _, article := range articles {
		// formate date from yyyy-mm-dd to yyyymmdd
		if date == strings.Replace(article.Date, "-", "", -1) {
			results = append(results, article)
		}
	}

	return results
}
