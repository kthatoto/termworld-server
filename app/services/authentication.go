package services

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) error {
	_, ok := c.Get("currentUser")
	if !ok {
		message := "Valid token is required"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}
	currentUser := CurrentUser(c)
	if !currentUser.Accepted {
		message := "The token is not accepted"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}
	return nil
}
