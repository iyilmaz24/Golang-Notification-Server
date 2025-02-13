package main

import (
	"fmt"
	"net/http"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/services"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	
	fmt.Fprintln(w, "Golang Monitoring Server Catch-All")
}

func (app *application) dailyAnalyticsReport(w http.ResponseWriter, r *http.Request) { // Used for daily reports, log to DB + send email
	var err = app.verifyPostRequest(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}

	analyticsObj, err := app.getDailyAnalyticsObject(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}
	var notiService services.NotificationService

	sendEmail := analyticsObj.NotificationEmail == true
	sendSMS := analyticsObj.NotificationSMS == true
	emailError, smsError := app.handleAnalyticsReport(w, sendEmail, sendSMS, analyticsObj, notiService) // send email and sms notifications

	loggingInfo := app.getAnalyticsReportLoggingInfo(analyticsObj)

	if emailError != nil || smsError != nil {
		app.handleEmailSmsError(w, err, emailError, smsError, loggingInfo, notiService) // checks if email or sms service is not working, alerts using the other method, logs the event to DB
		return
	}
	notiService.LogEventToDb(loggingInfo, "")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification successful"))
}

func (app *application) urgentNotification(w http.ResponseWriter, r *http.Request) { // Used for critical alerts, log to DB + send SMS & email
	err := app.verifyPostRequest(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}

	notiObj, err := app.getNotificationObject(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}
	var notiService services.NotificationService

	sendEmail := notiObj.NotificationEmail == true
	sendSMS := notiObj.NotificationSMS == true
	emailError, smsError := app.handleNotification(w, sendEmail, sendSMS, notiObj, notiService) // send email and sms notifications

	loggingInfo := app.getNotificationLoggingInfo(notiObj)

	if emailError != nil || smsError != nil {
		app.handleEmailSmsError(w, err, emailError, smsError, loggingInfo, notiService) // checks if email or sms service is not working, alerts using the other method, logs the event to DB
		return
	}
	notiService.LogEventToDb(loggingInfo, "")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification successful"))
}

func (app *application) onDemandNotification(w http.ResponseWriter, r *http.Request) { // Used for testing from outside the VPC / AWS environment, send SMS & email
	err := app.verifyPostRequest(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}

	notiObj, err := app.getNotificationObject(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}
	if notiObj.AccessSecret != config.LoadConfig().AdminPassword { // protection against unauthorized access
		app.clientError(w, http.StatusUnauthorized)
		return
	}
	var emailService services.EmailService
	var smsService services.SmsService
	emailError := error(nil)
	smsError := error(nil)

	if notiObj.NotificationEmail == true {
		emailError = emailService.SendEmailNotification(notiObj)
	}
	if notiObj.NotificationSMS == true {
		smsError = smsService.SendSmsNotification(notiObj)
	}

	loggingInfo := app.getNotificationLoggingInfo(notiObj)

	if emailError != nil || smsError != nil {
		app.handleEmailSmsError(w, emailError, smsError, loggingInfo, notiService) // checks if email or sms service is not working, alerts using the other method, logs the event to DB
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification successful"))
}

func (app *application) routineNotification(w http.ResponseWriter, r *http.Request) { // Used by Lambda when everything healthy, log to DB - dont send SMS or email
	err := app.verifyPostRequest(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}

	notiObj, err := app.getNotificationObject(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}

	emailErr, smsErr := error(nil), error(nil)

	if notiObj.NotificationEmail {
		emailService := services.EmailService{}
		emailErr = emailService.SendEmailNotification(notiObj)
	}

	if notiObj.NotificationSMS {
		smsService := services.SmsService{}
		smsErr = smsService.SendSmsNotification(notiObj)
	}

	if emailErr != nil || smsErr != nil {
		app.emailSmsSendError(w, emailErr, smsErr)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification successful"))
}
