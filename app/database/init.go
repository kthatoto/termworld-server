package middlewares

import (
	"context"
	"time"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	err := errors.New("")
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@mongo:27017"))
	if err != nil {
		log.Fatal(err)
		return
	}
}
