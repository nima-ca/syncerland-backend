package services

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func GetDialer() *gomail.Dialer {
	smtp := os.Getenv("SENDER_SMTP")
	smtpPort, _ := strconv.Atoi(os.Getenv("SENDER_SMTP_PORT"))
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")

	return gomail.NewDialer(smtp, int(smtpPort), senderEmail, senderPassword)
}

func SendOTP(to string, otp string) {
	// Create Message
	m := gomail.NewMessage()

	// Add sender and receiver
	m.SetHeader("From", "syncerland@gmail.com")
	m.SetHeader("To", to)

	// Add subject and body
	m.SetHeader("Subject", "Welcome To Syncerland")
	m.SetBody("text/html", fmt.Sprintf("<p>You can verify your account with this code: %s</p>", otp))

	// Get dialer and send the email
	err := GetDialer().DialAndSend(m)

	if err != nil {
		fmt.Println(err)
	}
}
