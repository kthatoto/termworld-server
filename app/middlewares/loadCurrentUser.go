package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/models"
)

func LoadCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Termworld-Token")
		if len(token) > 0 {
			var user models.User
			err := db.Database.Collection("users").FindOne(
				context.Background(),
				bson.M{"token": token},
			).Decode(&user)
			if err == nil {
				c.Set("currentUser", user)
			}
		}
		c.Next()
	}
}
