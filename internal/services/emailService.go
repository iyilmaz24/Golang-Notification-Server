package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
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

	err := sendEmail( emailRecipients, emailSubject, emailContent)

	if err != nil {
		smsService := SmsService{}
		err = sendEmailFallback(emailRecipients, emailSubject, emailContent)
		var errorMessage string

		if err != nil {
			errorMessage = fmt.Sprintf("Error sending email with fallback & primary method for '%s'", emailSubject)
		} else {
			errorMessage = fmt.Sprintf("Error sending email with primary method for '%s', fallback method successful", emailSubject)
		}
		smsService.AlertEmailNotWorking(errorMessage)
		
		logger.GetLogger().ErrorLog.Print(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func (notiService *EmailService) AlertSmsNotWorking(attempts int, errorMessage string, errObj error) error {
	
	currentTime := time.Now()
	
	failureInfo := &failureInfo{
		FailedAttempts: attempts,
		ErrorTime: fmt.Sprintf("%02d:%02d", currentTime.Hour(), currentTime.Minute()),
		ErrorCode: "SMS-SEND-ERROR",
		ErrorMessage: fmt.Sprintf("%s:\n\n %v", errorMessage, errObj),
	}

	emailRecipients := []string{config.LoadConfig().GmailAddress}
	emailSubject := "URGENT: SMS Service Failure"
	emailContent := getSMSServiceFailureContent(*failureInfo)

	err := sendEmail(emailRecipients, emailSubject, emailContent)

	if err != nil {
		err = sendEmailFallback(emailRecipients, emailSubject, emailContent)
		var errorMessage string

		if err != nil {
			errorMessage = "Error sending email about SMS service failure with fallback & primary email methods"
		} else {
			errorMessage = "Error sending email about SMS service failure with primary email method, fallback method successful"
		}

		logger.GetLogger().ErrorLog.Print(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}
