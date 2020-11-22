package models

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/forms"
)

type Player struct {
	Name string `bson:"name"`
}

type PlayerModel struct{}

func playerCollection() *mongo.Collection {
	return db.Database.Collection("players")
}

func (m PlayerModel) Create(form forms.PlayerCreateForm, currentUser User) (httpStatus int, err error) {
	if len(form.Name) == 0 {
		return http.StatusBadRequest, errors.New("Name is required")
	}

	userPlayersCount, err := playerCollection().CountDocuments(
		context.Background(),
		bson.M{"userID": currentUser.ID},
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if int(userPlayersCount) >= currentUser.MaxPlayerCount {
		return http.StatusForbidden, errors.New("Your player count already reached max count")
	}

	count, err := playerCollection().CountDocuments(
		context.Background(),
		bson.M{"name": form.Name},
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count > 0 {
		return http.StatusConflict, errors.New("The name is already used")
	}

	_, err = playerCollection().InsertOne(
		context.Background(),
		bson.M{"name": form.Name, "userID": currentUser.ID},
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
