package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goware/emailx"
	db "github.com/kthatoto/termworld-server/app/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type loginNewRequestJson struct {
	Email string `json:"email"`
}

func generateTokenString() {
	return "aaaabbbbcccc"
}

func LoginNew(c *gin.Context) {
	var data loginNewRequestJson
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	if err := emailx.Validate(data.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	collection := db.Database.Collection("users")
	res, err := collection.UpdateOne(
		context.Background(),
		bson.M{ "email": data.Email },
		bson.M{ "token": generateTokenString() },
		mongo.UpdateOptions{ Upsert: true },
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}
}
