package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/services"
)

func (app *application) verifyPostRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return fmt.Errorf("invalid request method: %s", r.Method)
	}
	return nil
}

func (app *application) getDailyAnalyticsObject(w http.ResponseWriter, r *http.Request) (models.DailyAnalytics, error) { // marshalls and returns the dailyAnalytics object from the POST request body
	var analyticsObj models.DailyAnalytics
	err := json.NewDecoder(r.Body).Decode(&analyticsObj)

	if err != nil || analyticsObj.NotificationSubject == "" || analyticsObj.NotificationMessage == "" {
		app.clientError(w, http.StatusBadRequest)
		return models.DailyAnalytics{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return analyticsObj, nil
}

func (app *application) getNotificationObject(w http.ResponseWriter, r *http.Request) (models.Notification, error) { // marshalls and returns the notification object from the POST request body
	var notificationObj models.Notification
	err := json.NewDecoder(r.Body).Decode(&notificationObj)

	if err != nil || notificationObj.NotificationSubject == "" || notificationObj.NotificationMessage == "" {
		app.clientError(w, http.StatusBadRequest)
		return models.Notification{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return notificationObj, nil
}

func (app *application) handleNotification(w http.ResponseWriter, sendEmail bool, sendSMS bool, notiObj models.Notification, notiService services.NotificationService) (error, error) {
	var emailError error = nil
	var smsError error = nil

	if sendEmail {
		emailError = notiService.SendEmailNotification(notiObj)
	}
	if sendSMS {
		smsError = notiService.SendSmsNotification(notiObj)
	}
	return emailError, smsError
}

func (app *application) handleAnalyticsReport(w http.ResponseWriter, sendEmail bool, sendSMS bool, analyticsObj models.DailyAnalytics, notiService services.NotificationService) (error, error) {
	var emailError error = nil
	var smsError error = nil

	if sendEmail {
		emailError = notiService.SendEmailReport(analyticsObj)
	}
	if sendSMS {
		smsError = notiService.SendSmsReport(analyticsObj)
	}
	return emailError, smsError
}

func (app *application) handleEmailSmsError(w http.ResponseWriter, err error, emailError error, smsError error, loggingInfo *models.LoggingInfo, notiService services.NotificationService) {

	switch { // check if email or sms service is not working, alert using the other method, log the event to DB
	case emailError != nil && smsError == nil:
		notiService.AlertEmailNotWorking()
		notiService.LogEventToDb(loggingInfo, "Email service is not working")
	case smsError != nil && emailError == nil:
		notiService.AlertSmsNotWorking()
		notiService.LogEventToDb(loggingInfo, "SMS service is not working")
	case emailError != nil && smsError != nil:
		notiService.LogEventToDb(loggingInfo, "Both Email and SMS services are not working")
	}
	app.emailSmsSendError(w, err, emailError != nil, smsError != nil) // last 2 arguments convert to boolean
	return
}

func (app *application) getAnalyticsReportLoggingInfo(analyticsObj models.DailyAnalytics) *models.LoggingInfo {
	return &models.LoggingInfo{
		NotificationType:      analyticsObj.NotificationType,
		NotificationSource:    analyticsObj.NotificationSource,
		NotificationRecipient: analyticsObj.NotificationRecipient,
		NotificationTime:      analyticsObj.NotificationTime,
		NotificationDate:      analyticsObj.NotificationDate,
		NotificationTimezone:  analyticsObj.NotificationTimezone,
		NotificationSubject:   analyticsObj.NotificationSubject,
	}
}

func (app *application) getNotificationLoggingInfo(notificationObj models.Notification) *models.LoggingInfo {
	return &models.LoggingInfo{
		NotificationType:      notificationObj.NotificationType,
		NotificationSource:    notificationObj.NotificationSource,
		NotificationRecipient: notificationObj.NotificationRecipient,
		NotificationTime:      notificationObj.NotificationTime,
		NotificationDate:      notificationObj.NotificationDate,
		NotificationTimezone:  notificationObj.NotificationTimezone,
		NotificationSubject:   notificationObj.NotificationSubject,
	}
}
