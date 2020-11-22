package models

import (
	"context"
	"net/http"
	"net/smtp"
	"os"
	"fmt"

	"github.com/goware/emailx"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kthatoto/termworld-server/app/forms"
	db "github.com/kthatoto/termworld-server/app/database"
	"github.com/kthatoto/termworld-server/app/utils"
)

type User struct {
	Email string
	Token string
	Accepted bool
}

type UserModel struct{}

func (m UserModel) LoginNew(form forms.LoginForm) (httpStatus int, err error) {
	if err := emailx.Validate(form.Email); err != nil {
		return http.StatusBadRequest, err
	}

	upsert := true
	loginToken := utils.RandomString(12)
	_, err = db.Database.Collection("users").UpdateOne(
		context.Background(),
		bson.M{ "email": form.Email },
		bson.M{ "$set": bson.M{ "token": loginToken, "accepted": false } },
		&options.UpdateOptions{ Upsert: &upsert },
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = loginMailSend(form.Email, loginToken)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
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
