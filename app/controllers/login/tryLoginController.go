package login

import (
	"context"
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goware/emailx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "github.com/kthatoto/termworld-server/app/database"
)

type tryLoginRequestJson struct {
	Email string `json:"email"`
}

type UserFromDB struct {
	Email string
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
		c.JSON(http.StatusBadRequest, gin.H{ "error": "The email is already accepted" })
		return
	}

	matchStage := bson.D{{"$match", bson.D{{"operationType", "update"}}}}
	opts := options.ChangeStream().SetMaxAwaitTime(15 * time.Second)
	changeStream, err := userCollection.Watch(context.Background(), mongo.Pipeline{matchStage}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}
	for changeStream.Next(context.Background()) {
		fmt.Println(changeStream.Current)
	}
	fmt.Println("finished")
}
