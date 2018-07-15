package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Types
type Tag struct {
	Name string `json:"name,omitempty"`
}
type Article struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Date  string `json:"date,omitempty"`
	Body  string `json:"body,omitempty"`
	Tags  []*Tag `json:"tags,omitempty"`
}
type TagSearchResponse struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []string `json:"articles"`
	RelatedTags []*Tag   `json:"related_tags"`
}

var articles []Article

// Display all articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

// Display a single article
func GetArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(GetArticleHelper(articles, params["id"]))
}

// Display articles that are on that date and match the tag
func GetArticlesByDateAndTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(GetArticlesByDateAndTagHelper(articles, params["date"], params["tagName"]))
}

// Main controller
func main() {
	port := "8000"
	router := mux.NewRouter()

	// Mock data
	articles = append(articles, Article{
		Id:    "1",
		Title: "latest science shows that potato chips are better for you than sugar",
		Date:  "2016-09-22",
		Body:  "some text, potentially containing simple markup about how potato chips are great",
		Tags: []*Tag{
			&Tag{Name: "health"},
			&Tag{Name: "fitness"},
			&Tag{Name: "science"},
		},
	})
	articles = append(articles, Article{
		Id:    "2",
		Title: "stores struggle with rising demand on potato chips",
		Date:  "2016-09-22",
		Body:  "some more text",
		Tags: []*Tag{
			&Tag{Name: "finance"},
			&Tag{Name: "lifestyle"},
		},
	})
	articles = append(articles, Article{
		Id:    "3",
		Title: "farmers planting more potatoes",
		Date:  "2016-09-22",
		Body:  "some more text",
		Tags: []*Tag{
			&Tag{Name: "health"},
			&Tag{Name: "lifestyle"},
		},
	})

	router.HandleFunc("/articles", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticle).Methods("GET")
	router.HandleFunc("/tags/{tagName}/{date}", GetArticlesByDateAndTag).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
	log.Print("Started running on port " + port)
}
