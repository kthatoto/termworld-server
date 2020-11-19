package services

import (
	"net/smtp"
	"os"
	"fmt"
)

func LoginMailSend(to string, token string) error {
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
