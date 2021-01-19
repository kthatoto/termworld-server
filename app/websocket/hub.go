package websocket

import "fmt"

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			fmt.Printf("registered! current clients: %d\n", len(h.clients))
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}
			fmt.Printf("unregistered! current clients: %d\n", len(h.clients))
		// case message := <-h.broadcast:
		// 	for client := range h.clients {
		// 		select {
		// 		case client.send <- message:
		// 		default:
		// 			close(client.send)
		// 			delete(h.clients, client)
		// 		}
		// 	}
		}
	}
}
