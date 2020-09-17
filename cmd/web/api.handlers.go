package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type tournamentResponse struct {
	ID                    int       `json:"id"`
	Title                 string    `json:"title"`
	Location              string    `json:"location"`
	TournamentDate        time.Time `json:"tournamentDate"`
	MatchTimeStart        string    `json:"matchTimeStart"`
	MatchTimeEnd          string    `json:"matchTimeEnd"`
	AdditionalInformation string    `json:"additionalInformation"`
	IsOnline              bool      `json:"isOnline"`
	TimeControl           string    `json:"timeControl"`
	TournamentType        string    `json:"tournamentType"`
	Rated                 bool      `json:"rated"`
	Poster                string    `json:"poster"`
	TournamentContact     string    `json:"tournamentContact"`
	Created               time.Time `json:"created"`
	Expires               time.Time `json:"expires"`
}

type userResponse struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	Club             string    `json:"club"`
	EloStandard      string    `json:"eloStandard"`
	EloRapid         string    `json:"eloRapid"`
	LichessUsername  string    `json:"lichessUsername"`
	ChesscomUsername string    `json:"chesscomUsername"`
	Created          time.Time `json:"created"`
	Active           bool      `json:"active"`
}

func (app *application) getLatestTournaments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")

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

func (app *application) getSingleTournament(w http.ResponseWriter, r *http.Request) {
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

func (app *application) getUserData(w http.ResponseWriter, r *http.Request) {
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
