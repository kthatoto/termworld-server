package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"

	"github.com/kthatoto/termworld-server/app/services"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ConnectWith(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		currentUser := services.CurrentUser(c)
		client := &Client{conn: conn, currentUser: &currentUser}
		hub.register <- client
		go client.handleMessages(hub)
	}
}
