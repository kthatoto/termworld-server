package middlewares

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:rootroot@localhost:27017"))
	if err != nil {
		log.Fatal(err)
		return
	}
	Database = client.Database("tw")
}
