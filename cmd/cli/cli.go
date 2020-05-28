package cli

import "flag"

// ServerConfig defines the fields for our server configuration
type ServerConfig struct {
	Addr        string
	StaticFiles string
}

// NewServerConfig is used at the beginning of a server start up and returns the ServerConfig struct with string values of the config
func NewServerConfig() ServerConfig {
	cfg := new(ServerConfig)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP netword address")
	flag.StringVar(&cfg.StaticFiles, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()
	return *cfg
}
