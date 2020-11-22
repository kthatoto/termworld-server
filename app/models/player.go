package models

import (
	"errors"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kthatoto/termworld-server/app/forms"
	db "github.com/kthatoto/termworld-server/app/database"
)

type Player struct {
	Name string
}

type PlayerModel struct{}

func playerCollection() *mongo.Collection {
	return db.Database.Collection("player")
}

func (m PlayerModel) Create(form forms.PlayerCreateForm) (httpStatus int, err error){
	if len(form.Name) == 0 {
		return http.StatusBadRequest, errors.New("Name is required")
	}

	count, err := playerCollection().CountDocuments(
		context.Background(),
		bson.M{ "name": form.Name },
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count > 0 {
		return http.StatusBadRequest, errors.New("The name is already used")
	}
	_, err = playerCollection().InsertOne(
		context.Background(),
		bson.M{ "name": form.Name },
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}