package services

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/models"
)

func Authentication(c *gin.Context) error {
	user, ok := c.Get("currentUser")
	if !ok {
		message := "Valid token is required"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}
	if !user.(models.User).Accepted {
		message := "The token is not accepted"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}
	return nil
}
