package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/golangcollege/sessions"

	"github.com/ctfrancia/bcnchess/cmd/cli"
	"github.com/ctfrancia/bcnchess/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	serverConfig  *cli.ServerConfig
	staticFiles   string
	session       *sessions.Session
	tournaments   *mysql.TournamentModel
	templateCache map[string]*template.Template
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
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		staticFiles:   serverConfig.StaticFiles,
		session:       serverConfig.Session,
		tournaments:   &mysql.TournamentModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     serverConfig.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting on server: %s", serverConfig.Addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
