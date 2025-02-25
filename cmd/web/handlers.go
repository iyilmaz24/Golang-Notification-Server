package main

import (
	"fmt"
	"net/http"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/services"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) { // Catch-all for requests
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Golang Monitoring Server Catch-All")
}

func (app *application) dailyAnalyticsReport(w http.ResponseWriter, r *http.Request) { // Used for daily reports, log to DB + send SMS &/or email

	analyticsObj, err := app.getDailyAnalyticsObject(w, r)
	if err != nil {
		logger.GetLogger().ErrorLog.Println(err)
		return
	}

	emailErr, smsErr := error(nil), error(nil)

	if analyticsObj.NotificationEmail {
		emailService := services.EmailService{}
		emailErr = emailService.SendEmailReport(analyticsObj)
	}

	if analyticsObj.NotificationSMS {
		smsService := services.SmsService{}
		smsErr = smsService.SendSmsReport(analyticsObj)
	}

	if emailErr != nil || smsErr != nil {
		app.emailSmsSendError(w, emailErr, smsErr)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification successful"))
}

func (app *application) internalNotification(w http.ResponseWriter, r *http.Request) { // Used for internal alerts, log to DB + send SMS &/or email

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

func (app *application) adminNotification(w http.ResponseWriter, r *http.Request) { // Used for testing from outside the VPC / AWS environment, send SMS &/or email

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

func (app *application) cronNotification(w http.ResponseWriter, r *http.Request) { // Used by Lambda when everything healthy, log to DB, send SMS &/or email

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
