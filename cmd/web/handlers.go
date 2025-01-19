package main

import (
	"fmt"
	"net/http"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/services"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Golang Notification Server Catch-All")
	fmt.Fprintln(w, "Use Correct Routes and Methods.")
}

func (app *application) dailyAnalyticsReport(w http.ResponseWriter, r *http.Request) { // Used for daily reports, log to DB + send email
	var err = app.verifyPostRequest(w, r)
	if err != nil {
		app.errorLog.Println(err)
		return 
	}

	analyticsObj, err := app.getDailyAnalyticsObject(w, r) 
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// send daily analytics report by email notification
	// handle error if email service is not working
	
	// log the event to DB
	// return success message & status code

}

func (app *application) urgentNotification(w http.ResponseWriter, r *http.Request) { // Used for critical alerts, log to DB + send SMS & email
	err := app.verifyPostRequest(w, r)
	if err != nil {
		app.errorLog.Println(err)
		return 
	}

	notiObj, err := app.getNotificationObject(w, r)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	var notiService services.NotificationService;
	
	sendEmail := notiObj.NotificationEmail == true
	sendSMS := notiObj.NotificationSMS == true
	emailError, smsError := app.handleNotification(w, sendEmail, sendSMS, notiObj, notiService)

	if emailError != nil || smsError != nil {
		app.handleNotificationError(w, err, emailError, smsError, notiObj, notiService) // checks if email or sms service is not working, alerts using the other method, logs the event to DB
		return
	}

	// completed notification:
		// log the event to DB
		// return success message & status code

}

func (app *application) onDemandNotification(w http.ResponseWriter, r *http.Request) { // Used for testing from outside the VPC / AWS environment, log to DB + send SMS & email
	err := app.verifyPostRequest(w, r)
	if err != nil {
		app.errorLog.Println(err)
		return 
	}

	notiObj, err := app.getNotificationObject(w, r)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	if (notiObj.AccessSecret != config.LoadConfig().AccessSecret) { // protection against unauthorized access
		app.clientError(w, http.StatusUnauthorized)
		return
	}
	var notiService services.NotificationService;

	sendEmail := notiObj.NotificationEmail == true
	sendSMS := notiObj.NotificationSMS == true
	emailError, smsError := app.handleNotification(w, sendEmail, sendSMS, notiObj, notiService) // checks if email or sms service is not working, alerts using the other method, logs the event to DB

	if emailError != nil || smsError != nil {
		app.handleNotificationError(w, err, emailError, smsError, notiObj, notiService)
		return
	}

	// completed notification:
		// log the event to DB
		// return success message & status code

}


func (app *application) routineNotification(w http.ResponseWriter, r *http.Request) { // Used by Lambda when everything healthy, log to DB - dont send SMS or email
	err := app.verifyPostRequest(w, r)
	if err != nil {
		app.errorLog.Println(err)
		return 
	}

	notiObj, err := app.getNotificationObject(w, r)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	var notiService services.NotificationService;
	
	// healthy routine notification:
		// no need to send email or SMS
		// only log the notification event to DB
		// return success message & status code

}