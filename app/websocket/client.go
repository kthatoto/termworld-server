package websocket

import (
	"bytes"
	"log"
	"encoding/json"

	ws "github.com/gorilla/websocket"

	"github.com/kthatoto/termworld-server/app/models"
	"github.com/kthatoto/termworld-server/app/websocket/handlers"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	conn        *ws.Conn
	currentUser *models.User
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
				log.Printf("error: %v\n", err)
			}
			continue
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		var command handlers.Command
		if err = json.Unmarshal(message, &command); err != nil {
			log.Printf("error: %v\n", err)
			continue
		}
		var resp handlers.Response
		resp, err = handlers.Handle(client.currentUser, command)
		if err != nil {
			log.Println(err)
			continue
		}

		writer, err := client.conn.NextWriter(ws.TextMessage)
		if err != nil {
			log.Println(err)
			continue
		}
		respJson, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
			continue
		}
		writer.Write([]byte(respJson))
		writer.Close()
	}
}
