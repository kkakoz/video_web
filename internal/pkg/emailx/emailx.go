package emailx

import (
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"net/smtp"
	"net/textproto"
)

type option struct {
	From     string
	Password string
	Email    string
	Addr     string
}

var o = &option{}

func Init(viper *viper.Viper) error {
	err := viper.UnmarshalKey("email", o)
	if err != nil {
		return err
	}
	return nil
}

func Send(to string, subject string, html string) error {
	e := &email.Email{
		To:      []string{to},
		From:    o.From,
		Subject: subject,
		HTML:    []byte(html),
		Headers: textproto.MIMEHeader{},
	}
	return e.Send("smtp.163.com:25", smtp.PlainAuth("", o.Email, o.Password, "smtp.163.com"))
}
