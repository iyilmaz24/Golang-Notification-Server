package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)                                     // Catch-All
	mux.HandleFunc("/dailyAnalyticsReport", app.dailyAnalyticsReport) // Used for daily reports, log event to DB + send SMS &/or email
	mux.HandleFunc("/urgentNotification", app.urgentNotification)     // Used for critical alerts, log event to DB + send SMS &/or email
	mux.HandleFunc("/routineNotification", app.routineNotification)   // Used for successful routine health checks, log event to DB
	mux.HandleFunc("/onDemandNotification", app.onDemandNotification) // Used for testing from outside the VPC / AWS environment, send SMS &/or email - do not log to DB

	handler := app.enableCors(mux)

	return handler
}
