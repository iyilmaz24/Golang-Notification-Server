package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
)

func (app *application) enableCors(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			app.clientError(w, http.StatusMethodNotAllowed)
			logger.GetLogger().ErrorLog.Printf("(cors-middleware) Method not allowed: %s, %s", r.Method, r.URL.Path)
			return
		}
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

		_, ok := corsOrigin[origin]
		if !ok {
			app.clientError(w, http.StatusForbidden)                                               // respond with 403 Forbidden
			logger.GetLogger().ErrorLog.Printf("(cors-middleware) Origin not allowed: %s", origin) // log the origin that was not allowed
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", origin) // set origin in response header if allowed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		next.ServeHTTP(w, r)
	})
}
