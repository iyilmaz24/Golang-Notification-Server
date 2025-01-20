package services

import (
	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type NotificationService struct {

}

func (notiService *NotificationService) SendEmailReport(analyticsObj models.DailyAnalytics) error {

	// send daily analytics report by email notification

	return nil;
}

func (notiService *NotificationService) SendSmsReport(analyticsObj models.DailyAnalytics) error {

	// send daily analytics report by email notification

	return nil;
}

func (notiService *NotificationService) SendEmailNotification(notificationObj models.Notification) error {

	// send email notification

	return nil;
}

func (notiService *NotificationService) SendSmsNotification(notificationObj models.Notification) error {

	// send text notification

	return nil;
}

func (notiService *NotificationService) LogEventToDb(loggingInfo *models.LoggingInfo, errorString string) error {
	if errorString != "" {
		// log event to database
	} else {
		// log error string to database
	}

	return nil;
}

func (notiService *NotificationService) AlertSmsNotWorking() error {

	// alert when sms service is not working with email

	return nil;
}

func (notiService *NotificationService) AlertEmailNotWorking() error {
	
	// alert when email service is not working with sms

	return nil;
}

