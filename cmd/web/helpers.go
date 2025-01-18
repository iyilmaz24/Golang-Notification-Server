package main

import (
	"fmt"
	"net/http"
)

type notification struct {
	NotificationMethod string `json:"method"`
	NotificationUrgency string `json:"urgency"`
	NotificationRecipient string `json:"recipient"`
	NotificationStatus string `json:"status"`
	NotificationID string `json:"id"`
	NotificationType string `json:"type"`
	NotificationSource string `json:"source"`
	NotificationTime string `json:"time"`
	NotificationDate string `json:"date"`
	NotificationTimezone string `json:"timezone"`
	NotificationSubject string `json:"subject"`
	NotificationMessage string `json:"message"`
	AccessSecret string `json:"password"`
}

type dailyAnalytics struct {
	NotificationRecipient string `json:"recipient"`
	NotificationTime string `json:"time"`
	NotificationDate string `json:"date"`
	NotificationTimezone string `json:"timezone"`
	NotificationSubject string `json:"subject"`
	NotificationMessage string `json:"message"`
}

func (app *application) verifyPostRequest (w http.ResponseWriter, r *http.Request) (error) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return fmt.Errorf("invalid request method: %s", r.Method)
	}
	return nil
}

func (app *application) getDailyAnalyticsObject(w http.ResponseWriter, r *http.Request) (dailyAnalytics, error) {

	// marshall and return the dailyAnalytics object from the POST request body
	
}

func (app *application) getNotificationObject(w http.ResponseWriter, r *http.Request) (notification, error) {

	// marshall and return the notification object from the POST request body
	
}
