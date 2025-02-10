package services

import (
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

	err := sendEmail(notificationObj)

	if err != nil {
		smsService := SmsService{}
		err = sendEmailFallback(notificationObj)
		var errorMessage string

		if err != nil {
			errorMessage = fmt.Sprintf("Error sending email with fallback & primary method for '%s'", notificationObj.NotificationSubject)
			smsService.AlertEmailNotWorking(errorMessage)
		} else {
			errorMessage = fmt.Sprintf("Error sending email with primary method for '%s'", notificationObj.NotificationSubject)
			smsService.AlertEmailNotWorking(errorMessage)
		}

		return fmt.Errorf(errorMessage)
	}

	return nil
}

func (notiService *EmailService) AlertSmsNotWorking(errorMessage string) error {

	// alert when sms service is not working with email

	return nil
}
