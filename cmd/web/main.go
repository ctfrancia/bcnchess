package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ctfrancia/bcnchess/cmd/cli"
)

func main() {
	serverConfig := cli.NewServerConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/tournament/:id", showTournament)
	mux.HandleFunc("/tournament/create", createTournament)

	fileServer := http.FileServer(http.Dir(serverConfig.StaticFiles))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println(serverConfig)
	log.Printf("starting on server: %s", serverConfig.Addr)
	err := http.ListenAndServe(serverConfig.Addr, mux)

	log.Fatal(err)
}
