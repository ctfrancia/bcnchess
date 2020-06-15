package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *application) apiGetLatestTournaments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tournaments, err := app.tournaments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	jsonData, err := json.Marshal(tournaments)
	if err != nil {
		app.serverError(w, err)
	}
	w.Write(jsonData)
}

func (app *application) apiGetSingleTournament(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
	}

	t, err := app.tournaments.Get(id)
	if err != nil {
		app.notFound(w)
	}

	jsonData, err := json.Marshal(t)
	if err != nil {
		app.serverError(w, err)
	}

	w.Write(jsonData)
}

func (app *application) apiGetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uID := app.session.GetInt(r, "authenticatedUserID")
	u, err := app.users.Get(uID)
	if err != nil {
		app.notFound(w)
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		app.serverError(w, err)
	}

	w.Write(jsonData)
}
