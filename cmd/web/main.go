package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/config"
	"github.com/iyilmaz24/Golang-Notification-Server/internal/logger"
)

type application struct {
}

func main() {

	infoLog := log.New(os.Stdout, "***INFO LOG:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "***ERROR LOG:\t", log.Ldate|log.Ltime|log.Lshortfile)
	logger.InitLogger(infoLog, errorLog)

	app := &application{}

	appConfig := config.LoadConfig()

	srv := &http.Server{
		Addr:     appConfig.Port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	logger.GetLogger().InfoLog.Println("(cmd/web/main.go) starting server on", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Println(err)
	}

}
