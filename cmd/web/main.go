package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/tournament/:id", showTournament)
	mux.HandleFunc("/tournament/create", createTournament)

	log.Println("starting on server: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
