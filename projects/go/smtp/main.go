package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host       := os.Getenv("SMTP_HOST")
	port       := os.Getenv("SMTP_PORT")
	user       := os.Getenv("SMTP_USER")
	pass       := os.Getenv("SMTP_PASS")
	from_name  := os.Getenv("MAIL_FROM_NAME")
	from_email := os.Getenv("MAIL_FROM_EMAIL")
	to_name    := os.Getenv("MAIL_TO_NAME")
	to_email   := os.Getenv("MAIL_TO_EMAIL")
	subj       := os.Getenv("MAIL_SUBJ")
	body       := os.Getenv("MAIL_BODY")

	auth := smtp.PlainAuth("", user, pass, host)

    // NOTE: Using the backtick here ` works like a heredoc, which is why all the
    // rest of the lines are forced to the beginning of the line, otherwise the
    // formatting is wrong for the RFC 822 style
	message := `To: "` + to_name + `" <` + to_email + `>
From: "` + from_name + `" <` + from_email + `>
Subject: ` + subj + `

` + body + `
`

	err = smtp.SendMail(host+":"+port, auth, from_email, []string{to_email}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
