package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	const DIR = "./serving-static-files/public"
	fmt.Println("Server will start at http://localhost:8000/")

	var route = mux.NewRouter()
	var fs = http.FileServer(http.Dir(DIR))

	// route.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fs))
	route.PathPrefix("/public").Handler(http.StripPrefix("/public/", fs))

	log.Fatal(http.ListenAndServe(":8000", route))
}
