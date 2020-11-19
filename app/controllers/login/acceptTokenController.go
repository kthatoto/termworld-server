package login

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	db "github.com/kthatoto/termworld-server/app/database"
)

type User struct {
	Email string
	Token string
	Accepted bool
}

func AcceptToken(c *gin.Context) {
	token := c.Param("token")

	var user *User
	err := db.Database.Collection("users").FindOne(
		context.Background(),
		bson.M{ "token": token },
	).Decode(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{ "token": token })
}
