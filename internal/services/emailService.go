package services

import (
	"errors"
	"fmt"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type EmailService struct {
}

func (notiService *EmailService) SendEmailReport(analyticsObj models.DailyAnalytics) error {

	// send daily analytics report by email notification

	return nil
}

func (notiService *EmailService) SendEmailNotification(notificationObj models.Notification) error {

	emailContent := getEmailContent(notificationObj)
	emailSubject := notificationObj.NotificationSubject
	emailRecipients := notificationObj.NotificationRecipients

	err := sendEmail(emailSubject, emailRecipients, emailContent)

	if err != nil {
		smsService := SmsService{}
		err = sendEmailFallback(emailSubject, emailRecipients, emailContent)
		var errorMessage string

		if err != nil {
			errorMessage = fmt.Sprintf("Error sending email with fallback & primary method for '%s'", emailSubject)
			smsService.AlertEmailNotWorking(errorMessage)
		} else {
			errorMessage = fmt.Sprintf("Error sending email with primary method for '%s'", emailSubject)
			smsService.AlertEmailNotWorking(errorMessage)
		}

		return errors.New(errorMessage)
	}

	return nil
}

func (notiService *EmailService) AlertSmsNotWorking(errorMessage string) error {

	// alert when sms service is not working with email

	return nil
}
