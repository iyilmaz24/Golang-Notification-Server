package services

import (
	"fmt"
	"time"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
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

func (notiService *SmsService) AlertEmailNotWorking(fallbackStatus bool, errorMessage string, errObj error) error {

	if fallbackStatus == false { // if the email fallback method is not working
		err := sendText(config.LoadConfig().AlertPhoneNumbers, "URGENT: email primary & fallback method failure")
		if err != nil {
			logger.GetLogger().ErrorLog.Print("(services/smsService.go) email service total failure, alerting sms service failed")
		}
		return err
	}

	// send notification about the email service failure using the email fallback method

	currentTime := time.Now()

	failureInfo := &failureInfo{
		FailedAttempts: 2,
		ErrorTime:      fmt.Sprintf("%02d:%02d", currentTime.Hour(), currentTime.Minute()),
		ErrorCode:      "EMAIL-SEND-ERROR",
		ErrorMessage:   fmt.Sprintf("%s:\n\n %v", errorMessage, errObj),
	}

	err := handleEmailFailure(*failureInfo)
	return err
}
