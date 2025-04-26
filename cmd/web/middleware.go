package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

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

func (app *application) routeAndOriginMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		corsOrigin := config.LoadConfig().Cors // loads a map[string]bool of allowed origins

		origin := r.Header.Get("Origin")
		if origin == "" {
			referer := r.Header.Get("Referer")
			if referer != "" {
				if parsedURL, err := url.Parse(referer); err == nil {
					origin = fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host) // remove path from referer, only keep scheme and host
				} else {
					logger.GetLogger().ErrorLog.Printf("(cors-middleware) Error parsing referer: %s", err)
				}
			}
		}

		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "" {
			if subtle.ConstantTimeCompare([]byte(apiKey), []byte(config.LoadConfig().AdminPassword)) != 1 { // constant-time comparison to prevent timing attacks
				logger.GetLogger().ErrorLog.Printf("(middleware) Invalid API key for request to %s", r.URL.Path)
				app.clientError(w, http.StatusUnauthorized)
				return
			}
			logger.GetLogger().InfoLog.Printf("API key authenticated request from origin: %s", origin)
			setCorsHeaders(w, origin)
			next.ServeHTTP(w, r) // skip cors origin check if a valid API key is provided
			return
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

type IPRateLimiter struct {
	mu            sync.Mutex
	ipRequestsMap map[string][]time.Time // map of IP addresses to a list of request times
	requestLimit  int                    // max requests allowed in time window
	timeWindow    time.Duration          // time window for rate limiting
}

func NewIPRateLimiter(maxRequests int, window time.Duration) *IPRateLimiter {
	return &IPRateLimiter{
		ipRequestsMap: make(map[string][]time.Time),
		requestLimit:  maxRequests,
		timeWindow:    window,
	}
}

func (limiter *IPRateLimiter) rateLimitMiddleware(next http.Handler) http.Handler { // limits the number of requests from an IP address

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr
		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		now := time.Now()
		var recent []time.Time
		if requests, exists := limiter.ipRequestsMap[ip]; exists { // if IP exists in map
			for _, t := range requests {
				if now.Sub(t) < limiter.timeWindow { // if request is within the time window
					recent = append(recent, t) // add request to recent list
				}
			}
			limiter.ipRequestsMap[ip] = recent
		} else if !exists { // if IP does not exist in map, create a new entry
			limiter.ipRequestsMap[ip] = []time.Time{}
		}

		if len(recent) >= limiter.requestLimit { // if the number of requests in the time window exceeds the limit
			logger.GetLogger().ErrorLog.Printf("(rate-limit-middleware) Rate limit exceeded for IP %s on %s", ip, r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error":"Rate limit exceeded. Try again later."}`))
			return
		}
		limiter.ipRequestsMap[ip] = append(limiter.ipRequestsMap[ip], now) // add current request time to the list of requests for this IP

		next.ServeHTTP(w, r)
	})
}
