package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
)

var allowedRoutesAndMethods = map[string]struct{}{ // concatenated method and route as key
	"GET/":                      {}, // empty struct so that the compiler doesn't allocate memory for values
	"POST/dailyAnalyticsReport": {},
	"POST/internalNotification": {},
	"POST/cronNotification":     {},
	"POST/adminNotification":    {},
}

func setCorsHeaders(w http.ResponseWriter, origin string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func (app *application) validateRouteAndOrigin(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		corsOrigin := config.LoadConfig().Cors // loads a map[string]bool of allowed origins

		origin := r.Header.Get("Origin")
		if origin == "" {
			referer := r.Header.Get("Referer")
			if referer != "" {
				if parsedURL, err := url.Parse(referer); err == nil {
					origin = fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host) // remove path from referer, only keep scheme and host
				} else if err != nil {
					logger.GetLogger().ErrorLog.Printf("(cors-middleware) Error parsing referer: %s", err)
				}
			}
		}

		_, validOrigin := corsOrigin[origin]
		if !validOrigin {
			app.clientError(w, http.StatusForbidden)                                               // respond with 403 Forbidden
			logger.GetLogger().ErrorLog.Printf("(cors-middleware) Origin not allowed: %s", origin) // log the origin that was not allowed
			return
		}

		if r.Method == http.MethodOptions {
			setCorsHeaders(w, origin)
			w.WriteHeader(http.StatusOK)
			return
		}

		_, validRequest := allowedRoutesAndMethods[r.Method+r.URL.Path] // concatentate method and route
		if !validRequest {
			app.clientError(w, http.StatusMethodNotAllowed)
			logger.GetLogger().ErrorLog.Printf("(cors-middleware) Invalid route/method combination: %s, %s", r.Method, r.URL.Path)
			return
		}

		setCorsHeaders(w, origin)

		next.ServeHTTP(w, r)
	})
}
