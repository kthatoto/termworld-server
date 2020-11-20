package services

import (
	"context"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/models"
)

func Authentication(c *gin.Context) error {
	token := c.GetHeader("X-Termworld-Token")
	if len(token) == 0 {
		message := "Token is required"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}

	var user models.User
	err := db.Database.Collection("users").FindOne(
		context.Background(),
		bson.M{ "token": token },
	).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": err.Error() })
		return errors.New(err.Error())
	}
	if !user.Accepted {
		message := "The token is not accepted"
		c.JSON(http.StatusUnauthorized, gin.H{ "error": message })
		return errors.New(message)
	}
	return nil
}
