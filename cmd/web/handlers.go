package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Golang Notification Server Catch-All")
	fmt.Fprintln(w, "Use Correct Routes and Methods.")
}

// below dailyAnalyticsReport route should receive a POST request with a JSON payload
	// notification recipient: string indicating who the notification is for

	// notification time: string (HH:MM)
	// notification date: string (MM/DD/YYYY)
	// notification timezone: string (EST, CST, PST)

	// notification subject: string 1 sentence summary
	// notification message: string explanation

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

// below routes should receive a POST request with a JSON payload
	// notification methods: email, sms
	// notification urgency: high, medium, low
	// notification recipient: string indicating who the notification is for

	// notification status: string (pending, retrying)
	// notification id: string (unique identifier)
	// notification type: string indicating what type of notification it is (CRON, internal, testing)
	// notification source: string indicating where the notification came from (specific server, database, website)

	// notification time: string (HH:MM)
	// notification date: string (MM/DD/YYYY)
	// notification timezone: string (EST, CST, PST)

	// notification subject: string 1 sentence summary
	// notification message: string explanation

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

	// should have protection against unauthorized use (secret key, etc.)

}
