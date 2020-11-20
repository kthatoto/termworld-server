package login

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goware/emailx"
	"go.mongodb.org/mongo-driver/bson"

	db "github.com/kthatoto/termworld-server/app/database"
)

type tryLoginRequestJson struct {
	Email string `json:"email"`
}

type UserFromDB struct {
	Email string
	Token string
	Accepted bool
}

func TryLogin(c *gin.Context) {
	var data loginNewRequestJson
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	if err := emailx.Validate(data.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	userCollection := db.Database.Collection("users")
	var user UserFromDB
	err := userCollection.FindOne(
		context.Background(),
		bson.M{ "email": data.Email },
	).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	if user.Accepted {
		c.JSON(http.StatusOK, gin.H{ "token": user.Token })
		return
	}

	for i := 0; i < 15; i++ {
		userCollection.FindOne(
			context.Background(),
			bson.M{ "email": data.Email },
		).Decode(&user)
		if user.Accepted {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if !user.Accepted {
		c.Status(http.StatusContinue)
		return
	}
	c.JSON(http.StatusOK, gin.H{ "token": user.Token })
}
