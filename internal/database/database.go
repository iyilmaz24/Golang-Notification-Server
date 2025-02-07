package database

import "github.com/iyilmaz24/Golang-Notification-Server/internal/models"

type Repository struct {

}

func (repo *Repository) LogEventToDb(loggingInfo *models.LoggingInfo, errorString string) error {
	if errorString != "" {
		// log event to database
	} else {
		// log error string to database
	}

	return nil;
}