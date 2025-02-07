package services

import (
	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type EmailService struct {

}

func (notiService *EmailService) SendEmailReport(analyticsObj models.DailyAnalytics) error {

	// send daily analytics report by email notification

	return nil;
}

func (notiService *EmailService) SendEmailNotification(notificationObj models.Notification) error {

	// send email notification

	return nil;
}

func (notiService *EmailService) AlertSmsNotWorking() error {

	// alert when sms service is not working with email

	return nil;
}


