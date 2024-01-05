package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
	m.SetHeader("From", os.Getenv("SENDER_EMAIL"))
	m.SetHeader("To", to)

	// Add subject and body
	m.SetHeader("Subject", "Welcome To Syncerland")
	m.SetBody("text/html", fmt.Sprintf("<p>You can verify your account with this code: %s</p>", otp))

	// Get dialer and send the email
	maxRetries := 3
	retryInterval := time.Second * 5
	err := SendWithRetry(GetDialer().DialAndSend, m, maxRetries, retryInterval)

	if err != nil {
		fmt.Println(err)
	}
}

func SendWithRetry(sendFunc func(m ...*gomail.Message) error, msg *gomail.Message, maxRetries int, retryInterval time.Duration) error {
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err = sendFunc(msg)
		if err == nil {
			// Email sent successfully, break out of the loop
			break
		}

		fmt.Printf("Error sending email (attempt %d): %v\n", attempt, err)

		// Wait before retrying
		time.Sleep(retryInterval)
	}

	return err
}
