package services

import (
	"errors"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
)

func handleEmailFailure(failureInfo failureInfo) error {

	emailRecipients := []string{config.LoadConfig().GmailAddress}

	emailSubject := "URGENT: Email Service Failure"
	emailContent := getEmailServiceFailureContent(failureInfo)
	err := sendEmailFallback(emailRecipients, emailSubject, emailContent)

	if err != nil {
		logger.GetLogger().ErrorLog.Print("(services/helpers.go) email service total failure")
		err = sendText(config.LoadConfig().AlertPhoneNumbers, "URGENT: email primary & fallback method failure")

		if err != nil {
			logger.GetLogger().ErrorLog.Print("(services/helpers.go) email service total failure, alerting sms service failed")
			return errors.New("email service total failure, alerting sms service failed")
		}
		return err
	}

	logger.GetLogger().ErrorLog.Print("(services/helpers.go) error sending email, alert sent with email fallback method")
	return nil
}

func handleSmsFailure(failureInfo failureInfo) error {

	emailRecipients := []string{config.LoadConfig().GmailAddress}

	emailSubject := "URGENT: SMS Service Failure"
	emailContent := getSMSServiceFailureContent(failureInfo)
	err := sendEmail(emailRecipients, emailSubject, emailContent)

	if err != nil {
		err = sendEmailFallback(emailRecipients, emailSubject, emailContent)
		var errorMessage string

		if err != nil {
			errorMessage = "error sending email about SMS service failure, email service total failure"
			logger.GetLogger().ErrorLog.Printf("(services/helpers.go) %s", errorMessage)
			return errors.New(errorMessage)
		}

		errorMessage = "error sending email about SMS service failure with primary email method, fallback method successful"
		logger.GetLogger().ErrorLog.Printf("(services/helpers.go) %s", errorMessage)
	}

	return nil
}
