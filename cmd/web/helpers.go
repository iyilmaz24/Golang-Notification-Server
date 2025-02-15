package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

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

func (app *application) getAnalyticsReportLoggingInfo(analyticsObj models.DailyAnalytics) *models.LoggingInfo {
	return &models.LoggingInfo{
		NotificationType:       analyticsObj.NotificationType,
		NotificationSource:     analyticsObj.NotificationSource,
		NotificationRecipients: analyticsObj.NotificationRecipients,
		NotificationTime:       analyticsObj.NotificationTime,
		NotificationDate:       analyticsObj.NotificationDate,
		NotificationTimezone:   analyticsObj.NotificationTimezone,
		NotificationSubject:    analyticsObj.NotificationSubject,
	}
}

func (app *application) getNotificationLoggingInfo(notificationObj models.Notification) *models.LoggingInfo {
	return &models.LoggingInfo{
		NotificationType:       notificationObj.NotificationType,
		NotificationSource:     notificationObj.NotificationSource,
		NotificationRecipients: notificationObj.NotificationRecipients,
		NotificationTime:       notificationObj.NotificationTime,
		NotificationDate:       notificationObj.NotificationDate,
		NotificationTimezone:   notificationObj.NotificationTimezone,
		NotificationSubject:    notificationObj.NotificationSubject,
	}
}
