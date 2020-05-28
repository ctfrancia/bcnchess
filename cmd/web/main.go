package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ctfrancia/bcnchess/cmd/cli"
)

type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	serverConfig *cli.ServerConfig
}

func main() {
	serverConfig := cli.NewServerConfig()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		serverConfig: serverConfig,
	}

	srv := &http.Server{
		Addr:     serverConfig.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting on server: %s", serverConfig.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
