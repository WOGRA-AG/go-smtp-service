package smtpsender

import (
	"net/smtp"

	"log"
	"strings"

	"wogra.com/configReader"
)

func SmtpSender(from string, to []string, subject string, message string) error {

	conf := configReader.ReadSmtpConfiguration()

	smtpAddress := conf.Smtpserver + ":" + conf.Smtpport

	mailContent := "From: " + from + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		message + "\r\n"

	log.Print("Content: >>>>" + mailContent + "<<<<")
	msg := []byte(mailContent)

	auth := smtp.PlainAuth("", conf.User, conf.Password, conf.Smtpserver)
	err := smtp.SendMail(smtpAddress, auth, conf.SenderAddress, to, msg)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
