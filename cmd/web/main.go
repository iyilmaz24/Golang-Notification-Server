package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application {
		errorLog: errorLog,
		infoLog: infoLog,
	}

	appConfig := config.LoadConfig();

	srv := &http.Server {
		Addr: appConfig.Addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	
	infoLog.Println("Starting server on", srv.Addr);

	err := srv.ListenAndServe();
	if err != nil {
		errorLog.Println(err);
	}

}