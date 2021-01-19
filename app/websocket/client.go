package websocket

import (
	"bytes"
	"log"
	"time"

	ws "github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	conn *ws.Conn
	send chan []byte
}

func (client *Client) handleMessages(hub *Hub) {
	defer func() {
		hub.unregister <- client
		client.conn.Close()
	}()

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		writer, err := client.conn.NextWriter(ws.TextMessage)
		if err != nil {
			return
		}

		if string(message) == "requestMap" {
			writer.Write([]byte("response map!!!"))
		} else {
			writer.Write([]byte("not supported"))
		}
		writer.Close()
	}
}
