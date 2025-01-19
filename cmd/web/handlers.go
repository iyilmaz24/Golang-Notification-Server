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

}

func (app *application) routineNotification(w http.ResponseWriter, r *http.Request) { // Used by Lambda when everything healthy, log to DB
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

	if (notiObj.AccessSecret == "" || notiObj.AccessSecret != config.LoadConfig().AccessSecret) { // protection against unauthorized access
		app.clientError(w, http.StatusUnauthorized)
		return
	}
	var notiService services.NotificationService;

	err = notiService.SendEmailNotification(notiObj);
	if err != nil {
		app.errorLog.Println(err)
		app.emailNotificationError(w, err)
		return
	}

	err = notiService.SendTextNotification(notiObj);
	if err != nil {
		app.errorLog.Println(err)
		app.smsNotificationError(w, err)
		return
	}

}
