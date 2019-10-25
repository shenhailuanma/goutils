package mail

import (
	"gopkg.in/gomail.v2"
)

/**
doc: https://tools.ietf.org/html/rfc2822.html
*/

type Mailer struct {
	username string		// email sender's username
	password string		// email sender's password
	host     string		// email server host
	port     int		// email server port
}

type MailInfo struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewMailer(username, password, host string, port int) (Mailer, error)  {
	var mailer = Mailer{username:username, password:password, host:host, port:port}
	return mailer, nil
}


func (mailer *Mailer)MailSendText(mailInfo MailInfo) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", mailInfo.From)
	msg.SetHeader("To", mailInfo.To)
	msg.SetHeader("Subject", mailInfo.Subject)
	msg.SetBody("text/plain", mailInfo.Body)

	d := gomail.NewDialer(mailer.host, mailer.port, mailer.username, mailer.password)
	err := d.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}

func (mailer *Mailer)MailSendHtml(mailInfo MailInfo) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", mailInfo.From)
	msg.SetHeader("To", mailInfo.To)
	msg.SetHeader("Subject", mailInfo.Subject)
	msg.SetBody("text/html", mailInfo.Body)

	d := gomail.NewDialer(mailer.host, mailer.port, mailer.username, mailer.password)
	err := d.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}
