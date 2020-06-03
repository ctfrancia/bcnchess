package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	t, err := app.tournaments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Tournaments: t})
}

func (app *application) aboutPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.page.tmpl", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
