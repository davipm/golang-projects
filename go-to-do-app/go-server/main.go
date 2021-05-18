package main

import (
	"fmt"
	"go-server/router"
	"log"
	"net/http"
	//"go-to-do-app/go-server/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
