package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ctfrancia/bcnchess/cmd/cli"
	_ "github.com/go-sql-driver/mysql"
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

	db, err := openDB(serverConfig.Dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
