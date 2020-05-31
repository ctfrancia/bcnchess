package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golangcollege/sessions"
)

// todo: add conditional ssr flag for server side rendering. If true then user needs to provide -templateDir flag
// for where the files can be found.

// ServerConfig defines the fields for our server configuration
type ServerConfig struct {
	Addr           string
	Dsn            string
	StaticFiles    string
	Session        *sessions.Session
	secret         string
	secretLifetime string
}

// NewServerConfig is used at the beginning of a server start up and returns the ServerConfig struct with string values of the config
func NewServerConfig() *ServerConfig {
	cfg := new(ServerConfig)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP netword address")
	flag.StringVar(&cfg.StaticFiles, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.Dsn, "dsn", "chess-web:password", "MySQL data source name <user>:<password>")
	flag.StringVar(&cfg.secret, "secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.StringVar(&cfg.secretLifetime, "secretLifetime", "12", "Secret key's lifetime")

	cfg.Dsn += "@/bcnchess?parseTime=true"

	cfg.Session = sessions.New([]byte(cfg.secret))
	lt, err := strconv.Atoi(cfg.secretLifetime)
	if err != nil {
		fmt.Fprint(os.Stdout, "-secretLifetime's value must be of type int")
	}
	cfg.Session.Lifetime = time.Duration(lt) * time.Hour
	cfg.Session.Secure = true

	return cfg
}
