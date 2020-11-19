package login

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	db "github.com/kthatoto/termworld-server/app/database"
)

func AcceptToken(c *gin.Context) {
	token := c.Param("token")

	_, err := db.Database.Collection("users").UpdateOne(
		context.Background(),
		bson.M{ "token": token },
		bson.M{ "$set": bson.M{ "accepted": true } },
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{ "token": token })
}
