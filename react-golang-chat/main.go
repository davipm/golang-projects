package main

import (
	"fmt"
	"go-basics/react-golang-chat/websocket"
	"net/http"
)

func serverWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit")
	var conn, err = websocket.Upgrade(w, r)

	if err != nil {
		var _, _ = fmt.Fprintf(w, "%v\n", err)
	}

	var client = &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	var pool = websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})
}

func main() {
	setupRoutes()
	fmt.Println("server running at localhost::8000")
	_ = http.ListenAndServe(":8000", nil)
}
