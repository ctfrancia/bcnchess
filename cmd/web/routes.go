package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logger, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	mux.Get("/tournament/create", dynamicMiddleware.ThenFunc(app.createTournamentForm))
	mux.Post("/tournament/create", dynamicMiddleware.ThenFunc(app.createTournament))

	mux.Get("/tournament/:id", dynamicMiddleware.ThenFunc(app.showTournament))

	fileServer := http.FileServer(http.Dir(app.staticFiles))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
