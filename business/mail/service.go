package mail

import (
	"vidlearn-final-projcect/config"

	"gopkg.in/gomail.v2"
)

type Service interface {
	SendMail(mail *Mail) (*Mail, error)
}

type mailService struct {
	config *config.AppConfig
}

func CreateService(config *config.AppConfig) Service {
	return &mailService{
		config: config,
	}
}

func (service *mailService) SendMail(mail *Mail) (*Mail, error) {
	if mail.Type == "reset" {
		mail.Type = "Reset Password"
	}

	if mail.Type == "verify" {
		mail.Type = "Verify Email"
	}

	if mail.Type == "reset-success" {
		mail.Type = "Reset Password Successful"
	}

	from := service.config.Mail.Sender + " - " + mail.Type + " <" + service.config.Mail.Username + ">"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", mail.To)
	mailer.SetAddressHeader("Cc", mail.From, "")
	mailer.SetHeader("Subject", mail.Subject)
	mailer.SetBody("text/html", mail.Body)

	dialer := gomail.NewDialer(
		service.config.Mail.Host,
		service.config.Mail.Port,
		service.config.Mail.Username,
		service.config.Mail.Password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return nil, err
	}

	return mail, nil
}
