package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/dailyAnalyticsReport", app.dailyAnalyticsReport)
	mux.HandleFunc("/urgentNotification", app.urgentNotification)
	mux.HandleFunc("/routineNotification", app.routineNotification)
	mux.HandleFunc("/onDemandNotification", app.onDemandNotification)

	handler := app.enableCors(mux);

	return handler;
}