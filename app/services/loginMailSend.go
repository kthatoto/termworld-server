package services

import (
	// "net/smtp"
	"fmt"
	"os"
)

// func LoginMailSend(to string) {
func LoginMailSend() {
	fmt.Println(os.Getenv("EMAIL"))
	fmt.Println(os.Getenv("EMAIL_PASSWORD"))
}
