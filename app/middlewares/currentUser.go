package services

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/kthatoto/termworld-server/app/models"
	db "github.com/kthatoto/termworld-server/app/database"
)

var CurrentUser *models.User
var CurrentUserExists bool

func LoadCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Termworld-Token")
		if len(token) > 0 {
			err := db.Database.Collection("users").FindOne(
				context.Background(),
				bson.M{ "token": token },
			).Decode(CurrentUser)
			if err {
				CurrentUserExists = false
			}
		} else {
			CurrentUserExists = false
		}
		c.Next()
	}
}
