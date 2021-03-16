package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type thumbnailRequest struct {
	Url string `json:"url"`
}

type screenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
}

type screenshotAPIResponse struct {
	Screenshot string `json:"screenshot"`
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest
	err := json.NewDecoder(r.Body).Decode(&decoded)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var apiRequest = screenshotAPIRequest{
		Token:          os.Getenv("TOKEN"),
		Url:            decoded.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}

	jsonString, err := json.Marshal(apiRequest)
	checkError(err)

	req, err := http.NewRequest("POST", "https://screenshotapi.net/api/v1/screenshot", bytes.NewBuffer(jsonString))
	checkError(err)

	req.Header.Set("Content-Type", "application/json")
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	defer response.Body.Close()

	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)

	_, err = fmt.Fprintf(w, `{ "screenshot": "%s" }`, apiResponse.Screenshot)
	checkError(err)
}

func main() {
	const client = "./client/build"
	http.HandleFunc("/api/thumbnail", thumbnailHandler)

	fs := http.FileServer(http.Dir(client))
	http.Handle("/", fs)

	fmt.Println("Server listening on port: 8080")
	log.Panic(http.ListenAndServe(":8080", nil))
}
