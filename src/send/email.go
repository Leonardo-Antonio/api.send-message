package send

import (
	"net/smtp"

	"github.com/Leonardo-Antonio/api.send-message/src/env"
	"github.com/jordan-wright/email"
)

func Many(tpl string, to ...string) error {
	e := email.NewEmail()
	e.From = "Leonardo Antonio Nolasco Leyva <" + env.Credentials.Email + ">"
	e.To = to
	e.Subject = "Message"
	e.HTML = []byte(tpl)
	if err := e.Send("smtp.gmail.com:587", smtp.PlainAuth(
		"", env.Credentials.Email,
		env.Credentials.Password,
		"smtp.gmail.com",
	)); err != nil {
		return err
	}

	return nil
}
