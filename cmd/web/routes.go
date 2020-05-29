package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logger, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))

	mux.Get("/tournament/create", http.HandlerFunc(app.createTournamentForm))
	mux.Post("/tournament/create", http.HandlerFunc(app.createTournament))

	mux.Get("/tournament/:id", http.HandlerFunc(app.showTournament))

	fileServer := http.FileServer(http.Dir(app.serverConfig.StaticFiles))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
