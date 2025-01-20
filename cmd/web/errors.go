package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)


func (app *application) serverError (w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError (w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound (w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) emailSmsSendError (w http.ResponseWriter, err error, email bool, sms bool) {
	var errorMessage string
	switch {
		case email && sms:
			errorMessage = "Email and SMS notification error"
		case email:
			errorMessage = "Email notification error"
		case sms:
			errorMessage = "SMS notification error"
		default:
			errorMessage = "Error sending notification"
	}
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Printf("%s: %s\n%s", errorMessage, err.Error(), trace)

	http.Error(w, errorMessage, http.StatusInternalServerError)
}
