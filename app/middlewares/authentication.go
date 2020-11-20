package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kthatoto/termworld-server/app/services"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("currentUser")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "error": "Valid token is required" })
			return
		}
		currentUser := services.CurrentUser(c)
		if !currentUser.Accepted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "error": "The token is not accepted" })
			return
		}
	}
}
