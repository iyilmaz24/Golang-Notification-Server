package services

import (
	"fmt"
	"net/smtp"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
)

func sendEmail(toEmails []string, emailSubject string, emailContent []byte) error {
	password := config.LoadConfig().GmailAppPassword
	from := config.LoadConfig().GmailAddress
	smtpHost := config.LoadConfig().GmailSmtpHost

	if len(toEmails) == 0 {
		logger.GetLogger().ErrorLog.Printf("No email recipients found for '%s'", emailSubject)
		return fmt.Errorf("no email recipients found for '%s'", emailSubject)
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":587", auth, from, toEmails, emailContent)
	if err != nil {
		logger.GetLogger().ErrorLog.Printf("Error sending email for '%s'", emailSubject)
		return fmt.Errorf("error sending email for '%s'", emailSubject)
	}

	logger.GetLogger().InfoLog.Printf("Emails sent successfully for '%s'", emailSubject)
	return nil
}

func sendEmailFallback(toEmails []string, emailSubject string, emailContent []byte) error {
	return nil
}
