package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

var articles Articles

func returnAllArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(writer).Encode(articles)
}

func returnSingleArticle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: Single Articles Endpoint")
	vars := mux.Vars(request)
	key := vars["id"]

	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(writer).Encode(article)
		}
	}
}

func testPostArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Test POST endpoint worked")
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "HomePage Endpont Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	articles = Articles{
		Article{
			Id: "1",
			Title:   "Hello 1",
			Desc:    "Test Description",
			Content: "Hello, World",
		},
		Article{
			Id: "2",
			Title:   "Hello 2",
			Desc:    "Test Description",
			Content: "Hello, World",
		},
	}

	handleRequests()
}
