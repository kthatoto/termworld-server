package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Connect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Reader, nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}
