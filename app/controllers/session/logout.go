package session

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kthatoto/termworld-server/app/services"
	db "github.com/kthatoto/termworld-server/app/database"
)

func Logout(c *gin.Context) {
	currentUser := services.CurrentUser(c)

	_, err := db.Database.Collection("users").UpdateOne(
		context.Background(),
		bson.M{ "token": currentUser.Token },
		bson.M{ "$set": bson.M{ "token": nil, "accepted": false } },
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}

	c.Status(http.StatusOK)
}
