package mail

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewMail(from, to, subject, body string) *Mail {
	return &Mail{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}
}
