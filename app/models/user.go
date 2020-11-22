package models

import (
	"context"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/goware/emailx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/forms"
	"github.com/kthatoto/termworld-server/app/utils"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	Email          string             `bson:"email"`
	Token          string             `bson:"token"`
	Accepted       bool               `bson:"accepted"`
	MaxPlayerCount int                `bson:"maxPlayerCount"`
}

type UserModel struct{}

func userCollection() *mongo.Collection {
	return db.Database.Collection("users")
}

func (m UserModel) LoginNew(form forms.LoginForm) (httpStatus int, err error) {
	if err := emailx.Validate(form.Email); err != nil {
		return http.StatusBadRequest, err
	}

	loginToken := utils.RandomString(12)
	existed, err := checkExistenceByEmail(form.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if existed {
		_, err = userCollection().UpdateOne(
			context.Background(),
			bson.M{"email": form.Email},
			bson.M{"$set": bson.M{"token": loginToken, "accepted": false}},
		)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		_, err = userCollection().InsertOne(
			context.Background(),
			bson.M{
				"email":          form.Email,
				"maxPlayerCount": 1,
				"token":          loginToken,
				"accepted":       false,
			},
		)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	err = loginMailSend(form.Email, loginToken)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (m UserModel) TryLogin(form forms.LoginForm) (token string, httpStatus int, err error) {
	if err := emailx.Validate(form.Email); err != nil {
		return token, http.StatusBadRequest, err
	}

	var user User
	err = userCollection().FindOne(
		context.Background(),
		bson.M{"email": form.Email},
	).Decode(&user)
	if err != nil {
		return token, http.StatusBadRequest, err
	}

	if user.Accepted {
		return user.Token, http.StatusOK, nil
	}

	for i := 0; i < 10; i++ {
		userCollection().FindOne(
			context.Background(),
			bson.M{"email": form.Email},
		).Decode(&user)
		if user.Accepted {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if !user.Accepted {
		return token, http.StatusRequestTimeout, nil
	}
	return user.Token, http.StatusOK, nil
}

func loginMailSend(to string, token string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	loginLink := fmt.Sprintf("%s/login/%s", os.Getenv("HOST"), token)
	msg := []byte("" +
		"From: termworld <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: Welcome to termworld!\r\n" +
		"\r\n" +
		"Please click the following link to login\r\n" +
		loginLink +
		"")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
	return err
}

func checkExistenceByEmail(email string) (bool, error) {
	count, err := userCollection().CountDocuments(
		context.Background(),
		bson.M{"email": email},
	)
	if err != nil {
		return false, err
	}
	return (count > 0), nil
}
