package main

import (
	"net/http"
	"time"
)

func (app *application) routes() http.Handler {

	limiter := NewIPRateLimiter(15, time.Minute) // 15 requests per minute rate limiter

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)                                     // Catch-All
	mux.HandleFunc("/dailyAnalyticsReport", app.dailyAnalyticsReport) // Used for daily reports, log event to DB + send SMS &/or email
	mux.HandleFunc("/internalNotification", app.internalNotification) // Used for internal alerts, log event to DB + send SMS &/or email
	mux.HandleFunc("/cronNotification", app.cronNotification)         // Used for routine health checks, log event to DB
	mux.HandleFunc("/adminNotification", app.adminNotification)       // Used for testing from outside the VPC / AWS environment, send SMS &/or email - do not log to DB

	handler := app.routeAndOriginMiddleware(
		app.apiKeyMiddleware(
			limiter.rateLimitMiddleware(mux))) // in order of validate route and origin, validate API key, then rate limit requests

	return handler
}
