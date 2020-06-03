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
	// filePath := uploadFile(w, r)
	now := time.Now()
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "additionalInformation", "expires") // more will be added here
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}
	c := convBool(form.Get("isOnline"))
	fmt.Println(c)

	data := &models.Tournament{
		Title:                 form.Get("title"),
		Location:              "location",
		TournamentDate:        time.Now(),
		MatchTimeStart:        "matchTimeStart",
		MatchTimeEnd:          "match time end",
		AdditionalInformation: "additionalInformation",
		IsOnline:              true,
		TimeControl:           "timeControl",
		TournamentType:        "tournamentType",
		Rated:                 false,
		Poster:                "filePath",
		Created:               now,
		Expires:               now.AddDate(0, 0, 1),
		// TODO: above needs to be changed once I know the example string that we are receiving
		// https://stackoverflow.com/questions/25845172/parsing-date-string-in-go
	}
	id, err := app.tournaments.Insert(data)
	if err != nil {
		app.serverError(w, err)
	}
	app.session.Put(r, "flash", "Tournament successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/tournament/%d", id), http.StatusSeeOther)
}
