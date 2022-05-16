package mail

import (
	"fmt"
	"net/smtp"
	"vidlearn-final-projcect/config"
)

type Service interface {
	SendMail(mail *Mail) (interface{}, error)
}

type mailService struct {
	config *config.AppConfig
}

func CreateService(config *config.AppConfig) Service {
	return &mailService{
		config: config,
	}
}

func (service *mailService) SendMail(mail *Mail) (interface{}, error) {
	if mail.Type == "reset" {
		mail.Type = "Reset Password"
	}

	if mail.Type == "verify" {
		mail.Type = "Verify Email"
	}

	if mail.Type == "reset-success" {
		mail.Type = "Reset Password Successful"
	}

	body := "From: " + service.config.Mail.Sender + " - " + mail.Type + " <" + service.config.Mail.Username + ">" + "\n" +
		"To: " + mail.To + "\n" +
		"Cc: " + mail.From + "\n" +
		"Subject: " + mail.Subject + "\n\n" +
		mail.Body

	auth := smtp.PlainAuth("", service.config.Mail.Username, service.config.Mail.Password, service.config.Mail.Host)
	smtpAddr := fmt.Sprintf("%s:%d", service.config.Mail.Host, service.config.Mail.Port)

	err := smtp.SendMail(smtpAddr, auth, service.config.Mail.Username, []string{mail.To, mail.From}, []byte(body))
	if err != nil {
		return nil, err
	}

	return mail.To, nil
}
