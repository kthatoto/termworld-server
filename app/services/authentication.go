package services

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/middlewares"
)

func Authentication(c *gin.Context) error {
	if !middlewares.CurrentUserExists {
		message := "Token is required"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}
	if !middlewares.CurrentUserExists.Accepted {
		message := "The token is not accepted"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}
	return nil
}
