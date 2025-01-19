package services

import (
	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type NotificationService struct {

}

func (notiService *NotificationService) SendAnalyticsReport(analyticsObj models.DailyAnalytics) error {

	// send daily analytics report by email notification

	return nil;
}

func (notiService *NotificationService) SendEmailNotification(notificationObj models.Notification) error {

	// send email notification

	return nil;
}

func (notiService *NotificationService) SendTextNotification(notificationObj models.Notification) error {

	// send text notification

	return nil;
}