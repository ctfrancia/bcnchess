package main

import (
	"crypto/tls"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ctfrancia/bcnchess/pkg/models"

	"github.com/golangcollege/sessions"

	"github.com/ctfrancia/bcnchess/cmd/cli"
	"github.com/ctfrancia/bcnchess/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog         *log.Logger
	infoLog          *log.Logger
	serverConfig     *cli.ServerConfig
	staticFiles      string
	tournamentImages string
	userImages       string
	session          *sessions.Session
	tournaments      interface {
		Insert(*models.Tournament) (int, error)
		Get(int) (*models.Tournament, error)
		Latest() ([]*models.Tournament, error)
	}
	templateCache map[string]*template.Template
	users         interface {
		Insert(*models.User) error
		Authenticate(string, string) (int, error)
		Get(int) (*models.User, error)
		UpdatePassword(int, string) error
	}
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
		errorLog:         errorLog,
		infoLog:          infoLog,
		staticFiles:      serverConfig.StaticFiles,
		tournamentImages: "./ui/static/imgs/tournaments",
		userImages:       "./ui/static/imgs/users",
		session:          serverConfig.Session,
		tournaments:      &mysql.TournamentModel{DB: db},
		templateCache:    templateCache,
		users:            &mysql.UserModel{DB: db},
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         serverConfig.Addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
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
