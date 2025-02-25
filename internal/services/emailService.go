package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type EmailService struct {
}

func (notiService *EmailService) SendEmailReport(analyticsObj models.DailyAnalytics) error {

	// send daily analytics report by email notification

	return nil
}

func (notiService *EmailService) SendEmailNotification(notificationObj models.Notification) error {

	emailRecipients := notificationObj.NotificationRecipients
	emailSubject := notificationObj.NotificationSubject
	emailContent := getEmailContent(notificationObj)

	err := sendEmail(emailRecipients, emailSubject, emailContent)

	if err != nil {
		smsService := SmsService{}
		err = sendEmailFallback(emailRecipients, emailSubject, emailContent)
		var errorMessage string

		if err != nil {
			errorMessage = fmt.Sprintf("(services/emailService.go) error sending email with fallback & primary method for '%s'", emailSubject)
			logger.GetLogger().ErrorLog.Print(errorMessage)
			smsService.AlertEmailNotWorking(false, errorMessage, err) // false for fallbackStatus
			return errors.New(errorMessage)
		} else {
			errorMessage = fmt.Sprintf("(services/emailService.go) error sending email with primary method for '%s', fallback method successful", emailSubject)
			logger.GetLogger().ErrorLog.Print(errorMessage)
			smsService.AlertEmailNotWorking(true, errorMessage, err) // true for fallbackStatus, can use email fallback method for alert
		}
		return err
	}

	return nil
}

func (notiService *EmailService) AlertSmsNotWorking(attempts int, errorMessage string, errObj error) error {

	currentTime := time.Now()

	failureInfo := &failureInfo{
		FailedAttempts: attempts,
		ErrorTime:      fmt.Sprintf("%02d:%02d", currentTime.Hour(), currentTime.Minute()),
		ErrorCode:      "SMS-SEND-ERROR",
		ErrorMessage:   fmt.Sprintf("%s:\n\n %v", errorMessage, errObj),
	}

	err := handleSmsFailure(*failureInfo)
	return err
}
