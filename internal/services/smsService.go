package services

import (
	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type SmsService struct {
}

func (notiService *SmsService) SendSmsReport(analyticsObj models.DailyAnalytics) error {

	// send daily analytics report by email notification

	return nil
}

func (notiService *SmsService) SendSmsNotification(notificationObj models.Notification) error {

	// send text notification

	return nil
}

func (notiService *SmsService) AlertEmailNotWorking() error {

	// alert when email service is not working with sms

	return nil
}
