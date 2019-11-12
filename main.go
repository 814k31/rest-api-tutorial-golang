package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(writer http.ResponseWriter, request *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Hello, World",
		},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(writer).Encode(articles)
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
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
