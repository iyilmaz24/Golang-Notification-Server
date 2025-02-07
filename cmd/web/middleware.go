package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
)

func (app *application) enableCors(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			app.clientError(w, http.StatusMethodNotAllowed)
			app.errorLog.Printf("***ERROR (cors-middleware): Method not allowed: %s", r.Method)
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
					app.errorLog.Printf("***ERROR (cors-middleware): Error parsing referer: %s", err)
				}
			}
		}

		_, ok := corsOrigin[origin]
		if !ok {
			app.clientError(w, http.StatusForbidden)                                          // respond with 403 Forbidden
			app.errorLog.Printf("***ERROR (cors-middleware): Origin not allowed: %s", origin) // log the origin that was not allowed
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", origin) // set origin in response header if allowed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
