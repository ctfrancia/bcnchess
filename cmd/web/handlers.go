package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ctfrancia/bcnchess/pkg/forms"
	"github.com/ctfrancia/bcnchess/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	t, err := app.tournaments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Tournaments: t})
	// app.render(w, r, "home.page.tmpl", nil)
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

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("firstName", "email", "password", "lastName")
	form.MaxLength("firstName", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 6)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	u := &models.User{
		FirstName:        form.Get("firstName"),
		LastName:         form.Get("lastName"),
		Email:            form.Get("email"),
		Password:         []byte(form.Get("password")),
		Club:             "Congres C.E",
		EloStandard:      "1700",
		EloRapid:         "1700",
		LichessUsername:  "",
		ChesscomUsername: "",
	}
	err = app.users.Insert(u)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", models.ErrEmailAlreadyExists)
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.session.Put(r, "flash", "Your signup was successful. Please login")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)

	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.session.Put(r, "authenticatedUserID", id)
	http.Redirect(w, r, "/tournament/create", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) aboutPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.page.tmpl", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// ADD TO HELPERS BELOW!!!!!!
func uploadFile(w http.ResponseWriter, r *http.Request) string {
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("poster")
	if err != nil {
		fmt.Println("Error getting file", err)
		return ""
	}
	defer file.Close()

	f, err := os.Create("./pkg/imgs/tournaments/" + handler.Filename)
	if err != nil {
		fmt.Println("Error saving", err)
		return ""
	}
	defer f.Close()

	io.Copy(f, file)
	return "./pkg/imgs/tournaments" + handler.Filename
}

func convBool(val string) bool {
	if val == "on" {
		return true
	}
	return false
}
