package cli

import "flag"

// ServerConfig defines the fields for our server configuration
type ServerConfig struct {
	Addr string
}

// NewServerConfig is used at the beginning of a server start up and returns the pointer to the config
func NewServerConfig() ServerConfig {
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()
	return ServerConfig{
		Addr: *addr,
	}
}
