package models

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/forms"
)

type Player struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
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
		return http.StatusForbidden, errors.New("Your players count already reached max count")
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
		bson.M{
			"name": form.Name,
			"userID": currentUser.ID,
			"live": false,
			"status": bson.M{},
		},
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (m PlayerModel) Index(currentUser User) ([]Player, int, error) {
	cursor, err := playerCollection().Find(
		context.Background(),
		bson.M{"userID": currentUser.ID},
	)
	var players []Player
	if err != nil {
		return players, http.StatusInternalServerError, err
	}
	if err := cursor.All(context.Background(), &players); err != nil {
		return players, http.StatusInternalServerError, err
	}
	return players, http.StatusOK, nil
}

func (m PlayerModel) FindByName(currentUser User, name string) (Player, error) {
	var player Player
	err := playerCollection().FindOne(
		context.Background(),
		bson.M{"userID": currentUser.ID, "name": name},
	).Decode(&player)
	if err != nil {
		return player, err
	}
	return player, nil
}

func (m PlayerModel) UpdateLive(player *Player, flag bool) (error) {
	var updatedDocument bson.M
	err := playerCollection().FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": player.ID},
		bson.M{"$set": bson.M{"live": flag}},
	).Decode(&updatedDocument)
	if err != nil {
		return err
	}
	return nil
}
