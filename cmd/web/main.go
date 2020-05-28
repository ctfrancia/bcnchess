package main

import (
	"log"
	"net/http"

	"github.com/ctfrancia/Ajedrez/cmd/cli"
)

func main() {
	// addr := flag.String("addr", ":4000", "Http network address")
	// flag.Parse()
	serverConfig := cli.NewServerConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/tournament/:id", showTournament)
	mux.HandleFunc("/tournament/create", createTournament)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	/*
		log.Printf("starting on server: %s", *addr)
		err := http.ListenAndServe(*addr, mux)
	*/

	log.Printf("starting on server: %s", serverConfig.Addr)
	err := http.ListenAndServe(serverConfig.Addr, mux)

	log.Fatal(err)
}
