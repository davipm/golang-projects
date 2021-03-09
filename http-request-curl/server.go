package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getAllUsers() {
	var url = "https://jsonplaceholder.typicode.com/users"
	var req, _ = http.NewRequest("GET", url, nil)
	var res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)

	log.Println(string(body))
}

func getAllUsersWithHeader() {
	var url = "https://jsonplaceholder.typicode.com/users"
	var req, _ = http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	var res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)

	log.Println(string(body))
}

func handlePostRequest() {
	var url = "https://jsonplaceholder.typicode.com/users"
	var payload = strings.NewReader("id=1")

	var req, _ = http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	var res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)

	log.Println(string(body))
}

func main() {
	getAllUsers()
	getAllUsersWithHeader()
	handlePostRequest()
}
