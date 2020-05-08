package mailer

import (
	"fmt"
	"os"

	"github.com/emersion/go-sasl"
)

var (
	auth    *sasl.Client
	address *string
	from    *string
)

// ConnectToSMTP connects to the SMTP server and assigns a reference to the global auth
func ConnectToSMTP() {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpFrom := os.Getenv("SMTP_FROM")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	smtpAddress := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	plainAuth := sasl.NewPlainClient("", smtpUsername, smtpPassword)

	auth = &plainAuth
	address = &smtpAddress
	from = &smtpFrom
}
