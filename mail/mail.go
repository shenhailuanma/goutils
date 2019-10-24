package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

var MailUsername = ""
var MailHost = ""
var MailPort = 0
var MailPassword = ""

func MailSendText(address string, subject string,  data string) error {

	msg := gomail.NewMessage()
	var from = fmt.Sprintf("Notifiction <%s>", MailUsername)
	msg.SetHeader("From", from)
	msg.SetHeader("To", address)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", data)

	d := gomail.NewDialer(MailHost, MailPort,MailUsername, MailPassword)
	err := d.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}


func MailSendHtml(address string, subject string,  data string) error {

	msg := gomail.NewMessage()

	var from = fmt.Sprintf("Notifiction <%s>", MailUsername)
	msg.SetHeader("From", from)
	msg.SetHeader("To", address)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", data)

	d := gomail.NewDialer(MailHost, MailPort, MailUsername, MailPassword)
	err := d.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}
