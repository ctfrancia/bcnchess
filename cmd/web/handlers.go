package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ctfrancia/bcnchess/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	t, err := app.tournaments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Tournaments: t})
}
func (app *application) showTournament(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
	}
	t, err := app.tournaments.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "tournament.page.tmpl", &templateData{Tournament: t})
}

func (app *application) createTournament(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	data := models.Tournament{
		Title:                 "tournament 1",
		Location:              "Mom's basement",
		MatchTimeStart:        now,
		MatchTimeEnd:          now.AddDate(0, 0, 1),
		AdditionalInformation: "Here is the additional information about my mom's basement, it's cold and damp",
		IsOnline:              false,
		TimeControl:           "3+2",
		TournamentType:        "Swiss",
		Rated:                 true,
		Poster:                "./ui/static/img/logo.png",
		Created:               now,
		Expires:               now.AddDate(0, 0, 1),
	}
	id, err := app.tournaments.Insert(data)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/tournament/:%d", id), http.StatusSeeOther)
	w.Write([]byte("create tournament"))
}
