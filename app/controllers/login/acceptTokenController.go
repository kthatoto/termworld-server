package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AcceptToken(c *gin.Context) {
	token := c.Param("token")
	c.JSON(http.StatusOK, gin.H{ "token": token })
}
