package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/tournament/:id", app.showTournament)
	mux.HandleFunc("/tournament/create", app.createTournament)

	fileServer := http.FileServer(http.Dir(app.serverConfig.StaticFiles))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
