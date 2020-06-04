package main

import (
	"errors"
	"net/http"

	"github.com/ctfrancia/bcnchess/pkg/forms"
	"github.com/ctfrancia/bcnchess/pkg/models"
)

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
	form.Required("firstName", "email", "password", "lastName", "retypePassword")
	form.PasswordsMatch(form.Get("password"), form.Get("retypePassword"))
	form.MaxLength("firstName", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 6)
	form.MinLength("retypePassword", 6)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
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

func (app *application) userProfile(w http.ResponseWriter, r *http.Request) {
	uID := app.session.GetInt(r, "authenticatedUserID")
	u, err := app.users.Get(uID)
	if err != nil {
		app.serverError(w, err)
	}
	app.render(w, r, "user.page.tmpl", &templateData{UserProfile: u})
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
	path := app.session.PopString(r, "redirectPathAfterLogin")
	if path != "" {
		http.Redirect(w, r, path, http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/tournament/create", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) changePasswordForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "passwordChange.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) updatePassword(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)

	u, err := app.users.Get(app.session.GetInt(r, "authenticatedUserID"))
	if err != nil {
		app.serverError(w, err)
	}

	id, err := app.users.Authenticate(u.Email, form.Get("oldPassword"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("oldPassword", "incorrect password")
			app.render(w, r, "passwordChange.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	form.MinLength("password1", 6)
	form.MinLength("password2", 6)
	form.PasswordsMatch(form.Get("password1"), form.Get("password2"))

	if !form.Valid() {
		app.render(w, r, "passwordChange.page.tmpl", &templateData{Form: form})
		return
	}

	err = app.users.UpdatePassword(id, form.Get("password1"))
	if err != nil {
		app.serverError(w, err)
	}

	app.session.Put(r, "flash", "password updated successfully")
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}
