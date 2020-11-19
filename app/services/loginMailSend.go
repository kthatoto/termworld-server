package services

import (
	"net/smtp"
	"os"
)

func LoginMailSend(to string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	msg := []byte("" +
		"From: 送信した人 <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 件名 subject です\r\n" +
		"\r\n" +
		"テスト\r\n" +
	"")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
	return err
}
