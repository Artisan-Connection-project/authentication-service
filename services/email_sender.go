package services

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailSender interface {
	SendResetEmail(to, token string) error
}

type emailSenderImpl struct {
	smtpAddr    string
	smtpAuth    smtp.Auth
	fromAddress string
}

func NewEmailSender(smtpAddr, username, password, fromAddress string) EmailSender {
	return &emailSenderImpl{
		smtpAddr:    smtpAddr,
		smtpAuth:    smtp.PlainAuth("", username, password, smtpAddr),
		fromAddress: fromAddress,
	}
}

func (es *emailSenderImpl) SendResetEmail(to, token string) error {
	e := email.NewEmail()
	e.From = es.fromAddress
	e.To = []string{to}
	e.Subject = "Password Reset"
	e.Text = []byte("Use this token to reset your password: " + token)
	return e.Send(es.smtpAddr, es.smtpAuth)
}
