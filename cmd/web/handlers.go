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

func (app *application) dailyAnalyticsReport(w http.ResponseWriter, r *http.Request) {
	
}

func (app *application) urgentNotification(w http.ResponseWriter, r *http.Request) {

}

func (app *application) routineNotification(w http.ResponseWriter, r *http.Request) {
	
}

func (app *application) onDemandNotification(w http.ResponseWriter, r *http.Request) {
	
}
