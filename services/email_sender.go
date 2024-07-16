package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type EmailService interface {
	SendCode(email, code string) error
}

type emailServiceImpl struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

func NewEmailService(from, password, smtpHost, smtpPort string) EmailService {
	return &emailServiceImpl{
		from:     from,
		password: password,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
	}
}

func (e *emailServiceImpl) SendCode(email, code string) error {
	to := []string{email}
	auth := smtp.PlainAuth("", e.from, e.password, e.smtpHost)

	t, err := template.ParseFiles("template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Your verification code \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct{ Passwd string }{Passwd: code})

	return smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.from, to, body.Bytes())
}
