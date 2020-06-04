package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ctfrancia/bcnchess/pkg/forms"
	"github.com/ctfrancia/bcnchess/pkg/models"
)

func (app *application) showTournament(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

func (app *application) createTournamentForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createTournament(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "multipart/form-data")
	now := time.Now()
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "additionalInformation", "location", "matchTimeStart", "tournamentDate", "tournamentContact")
	form.MatchesPattern("tournamentContact", forms.EmailRX)
	form.MaxLength("title", 100)

	if !form.Valid() {
		fmt.Println("errors", &templateData{Form: form})
		return
	}
	td, err := time.Parse("2006-01-02", form.Get("tournamentDate"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	data := &models.Tournament{
		Title:                 form.Get("title"),
		Location:              form.Get("location"),
		TournamentDate:        td,
		MatchTimeStart:        form.Get("matchTimeStart"),
		MatchTimeEnd:          "match time end",
		AdditionalInformation: form.Get("additionalInformation"),
		IsOnline:              convBool(form.Get("isOnline")),
		TimeControl:           form.Get("timeControl"),
		TournamentType:        form.Get("tournamentType"),
		Rated:                 convBool(form.Get("isRated")),
		Poster:                uploadFile(w, r),
		TournamentContact:     form.Get("tournamentContact"),
		Created:               now,
		Expires:               td.AddDate(0, 0, 1),
	}
	id, err := app.tournaments.Insert(data)
	if err != nil {
		app.serverError(w, err)
	}
	app.session.Put(r, "flash", "Tournament successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/tournament/%d", id), http.StatusSeeOther)
}
