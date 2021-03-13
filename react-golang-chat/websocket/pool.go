package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("Size of connection pool:", len(p.Clients))

			for client := range p.Clients {
				fmt.Println(client)
				_ = client.Conn.WriteJSON(Message{Type: 1, Body: "New user has joined!"})
			}
			break
		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("Size of connection pool:", len(p.Clients))

			for client := range p.Clients {
				fmt.Println(client)
				_ = client.Conn.WriteJSON(Message{Type: 1, Body: "New user has left!"})
			}
			break
		case message := <-p.Broadcast:
			fmt.Println("Sending message to all client in the pool")

			for client := range p.Clients {
				var err = client.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			break
		}
	}
}
