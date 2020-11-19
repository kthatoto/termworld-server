package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goware/emailx"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/services"
)

type loginNewRequestJson struct {
	Email string `json:"email"`
}

func generateTokenString() string {
	return "newtoken"
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
	upsert := true
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{ "email": data.Email },
		bson.M{ "$set": bson.M{
			"token": generateTokenString(),
		}},
		&options.UpdateOptions{ Upsert: &upsert },
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}

	err = services.LoginMailSend(data.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}
}
