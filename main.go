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

var articles []Article

// Display all articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
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

	router.HandleFunc("/articles", GetArticles).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
	log.Print("Started running on port " + port)
}
