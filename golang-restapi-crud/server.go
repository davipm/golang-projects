package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Rest api")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":3000", router))
}
