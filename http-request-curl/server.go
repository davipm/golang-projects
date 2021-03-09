package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = "https://jsonplaceholder.typicode.com"

func getAllUsers() {
	var url = URL + "/users"
	var req, _ = http.NewRequest("GET", url, nil)
	var res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func getAllUsersWithHeader() {
	var url = URL + "/users"
	var req, _ = http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	var res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func handlePostRequest() {
	var url = "https://reqres.in/api/users"
	var payload = strings.NewReader("name=test&jab=teacher")

	var req, _ = http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	var res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func main() {
	getAllUsers()
	getAllUsersWithHeader()
	handlePostRequest()
}
