package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/ctfrancia/bcnchess/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	t, err := app.tournaments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Tournaments: t})
}
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
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createTournament(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	data := models.Tournament{
		Title:                 r.PostForm.Get("title"),
		Location:              "Mom's basement",
		MatchTimeStart:        now,
		MatchTimeEnd:          now.AddDate(0, 0, 1),
		AdditionalInformation: r.PostForm.Get("additionalInformation"),
		IsOnline:              false,
		TimeControl:           "3+2",
		TournamentType:        "Swiss",
		Rated:                 true,
		Poster:                "./ui/static/img/logo.png",
		Created:               now,
		Expires:               now.AddDate(0, 0, 1),
		// TODO: above needs to be changed once I know the example string that we are receiving
		// https://stackoverflow.com/questions/25845172/parsing-date-string-in-go
	}

	errors := make(map[string]string)
	if strings.TrimSpace(data.Title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(data.Title) > 100 {
		errors["title"] = "This field is too long(maximmum is 100 characters)"
	}

	if strings.TrimSpace(data.AdditionalInformation) == "" {
		errors["additionalInformation"] = "This field cannot be blank"
	}
	// TODO: add here the part for checking the expiring time is entered

	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	id, err := app.tournaments.Insert(data)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/tournament/%d", id), http.StatusSeeOther)
}
