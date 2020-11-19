package controllers

import (
	"context"
	"log"
	"net/http"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginNew(c *gin.Context) {
	collection := db.Database.Collection("users")

	res, err := collection.InsertOne(context.Background(), bson.M{
		"Email": "aaaa",
		"Token": "cccc",
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": res.InsertedID,
		"status": "ok",
	})
}
