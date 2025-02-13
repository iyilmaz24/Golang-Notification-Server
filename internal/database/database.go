package database

import (
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

type Repository struct {
}

func (repo *Repository) LogEventToDb(loggingInfo *models.LoggingInfo, errorString string) error {
	if errorString == "" {
		// logger.GetLogger().InfoLog.Printf("(internal/database/database.go) Event logged to database '%s' @ %s", loggingInfo.NotificationSubject, loggingInfo.NotificationTime)
		logger.GetLogger().InfoLog.Printf("(internal/database/database.go) Need to add logic for logging notifications to database")
	} else {
		// logger.GetLogger().InfoLog.Printf("(internal/database/database.go) Error logged to database '%s'", errorString)
		logger.GetLogger().InfoLog.Printf("(internal/database/database.go) Need to add logic for logging errors to database")
	}

	return nil
}
