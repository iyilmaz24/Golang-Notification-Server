package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	logger.GetLogger().ErrorLog.Println(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) emailSmsSendError(w http.ResponseWriter, email error, sms error) {
	var errorMessage string

	switch {
	case email != nil && sms != nil:
		errorMessage = "email and sms notification error"
	case email != nil:
		errorMessage = "email notification error"
	case sms != nil:
		errorMessage = "sms notification error"
	}

	logger.GetLogger().ErrorLog.Print(errorMessage)
	http.Error(w, errorMessage, http.StatusInternalServerError) // respond to client with error message
}
