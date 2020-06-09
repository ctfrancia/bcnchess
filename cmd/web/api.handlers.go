package main

import (
	"encoding/json"
	"net/http"
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

func (app *application) getLatest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonData []byte

	tournaments, err := app.tournaments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	jsonData, err = json.Marshal(tournaments)
	if err != nil {
		app.serverError(w, err)
	}
	w.Write(jsonData)
}
