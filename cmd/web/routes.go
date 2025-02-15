package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)                                     // Catch-All
	mux.HandleFunc("/dailyAnalyticsReport", app.dailyAnalyticsReport) // Used for daily reports, log event to DB + send SMS &/or email
	mux.HandleFunc("/internalNotification", app.internalNotification) // Used for internal alerts, log event to DB + send SMS &/or email
	mux.HandleFunc("/cronNotification", app.cronNotification)         // Used for routine health checks, log event to DB
	mux.HandleFunc("/adminNotification", app.adminNotification)       // Used for testing from outside the VPC / AWS environment, send SMS &/or email - do not log to DB

	handler := app.validateRouteAndOrigin(mux)

	return handler
}
