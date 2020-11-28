package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ctfrancia/bcnchess/pkg/forms"
	"github.com/ctfrancia/bcnchess/pkg/models"
)

type errorResponse struct {
	Reason string `json:"reason"`
}

func (app *application) apiRegisterNewUser(w http.ResponseWriter, r *http.Request) {
	nu := make(map[string]string)
	// nu := models.NewUserJSON{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &nu)
	if err != nil {
		fmt.Println(err)
	}

	v := forms.NewValidator()
	v.ValidateFields(nu)
	v.MatchesPattern("email", forms.EmailRX)
	v.MaxLength("email", 35)
	v.MaxLength("firstName", 15)
	v.MaxLength("lastName", 25)
	v.MinLength("firsttName", 3)
	v.MinLength("lastName", 3)

	valid, errs := v.Valid()
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errs)
		_, _ = w.Write(b)
		return
	}

	u := &models.User{
		FirstName:        form.Get("firstName"),
		LastName:         form.Get("lastName"),
		Email:            form.Get("email"),
		Password:         []byte(form.Get("password")),
		Club:             form.Get("club"),
		EloStandard:      form.Get("eloStandard"),
		EloRapid:         form.Get("eloRapid"),
		LichessUsername:  form.Get("lichessUsername"),
		ChesscomUsername: form.Get("chesscomUsername"),
	}
	fmt.Printf("%#v: ", v)
	/*
		u := &models.User{
			FirstName:        form.Get("firstName"),
			LastName:         form.Get("lastName"),
			Email:            form.Get("email"),
			Password:         []byte(form.Get("password")),
			Club:             form.Get("club"),
			EloStandard:      form.Get("eloStandard"),
			EloRapid:         form.Get("eloRapid"),
			LichessUsername:  form.Get("lichessUsername"),
			ChesscomUsername: form.Get("chesscomUsername"),
		}

		err = app.users.Insert(u)
		if err != nil {
			if errors.Is(err, models.ErrDuplicateEmail) {
				estruct := errorResponse{
					Reason: "email exists",
				}

				er, err := json.Marshal(estruct)
				if err != nil {
					panic(err)
				}

				w.WriteHeader(http.StatusBadRequest)
				w.Write(er)
			} else {
				app.serverError(w, err)
			}
			return
		}
	*/
	w.WriteHeader(http.StatusOK)
}

func (app *application) apiGetLatestTournaments(w http.ResponseWriter, r *http.Request) {
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
