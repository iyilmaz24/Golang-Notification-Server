package logger

import (
	"log"
	"sync"
)

var (
	once     sync.Once
	instance *Logger
)

type Logger struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func InitLogger(infoLog *log.Logger, errorLog *log.Logger) {
	once.Do(func() {
		instance = &Logger{
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		}
	})
}

func GetLogger() *Logger {
	return instance
}
