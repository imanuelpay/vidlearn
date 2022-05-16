package mail

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
	Type    string
}

func NewMail(From, To, Subject, Body, Type string) *Mail {
	return &Mail{
		From:    From,
		To:      To,
		Subject: Subject,
		Body:    Body,
		Type:    Type,
	}
}
