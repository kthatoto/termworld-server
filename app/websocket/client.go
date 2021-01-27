package websocket

import (
	"bytes"
	"log"
	"encoding/json"

	ws "github.com/gorilla/websocket"

	"github.com/kthatoto/termworld-server/app/models"
	"github.com/kthatoto/termworld-server/game/command"
	"github.com/kthatoto/termworld-server/game/command/commands"
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
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		var cmd command.Command
		if err = json.Unmarshal(message, &cmd); err != nil {
			log.Printf("error: %v\n", err)
			continue
		}

		writer, _ := client.conn.NextWriter(ws.TextMessage)
		var resp commands.Response
		resp, err = command.Handle(client.currentUser, cmd)
		if err != nil {
			log.Println(err)
		}
		respJson, _ := json.Marshal(resp)
		writer.Write([]byte(respJson))
		writer.Close()
	}
}
