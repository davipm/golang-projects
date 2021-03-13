package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var conn, err = Upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return conn, err
	}

	return conn, nil
}
