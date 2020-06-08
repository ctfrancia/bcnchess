package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logger, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	mux.Get("/about", dynamicMiddleware.ThenFunc(app.aboutPage))

	mux.Get("/tournament/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createTournamentForm))
	mux.Post("/tournament/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createTournament))
	mux.Get("/tournament/:id", dynamicMiddleware.ThenFunc(app.showTournament))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))
	mux.Get("/user/profile", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.userProfile))
	mux.Get("/user/change-password", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.changePasswordForm))
	mux.Post("/user/change-password", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.updatePassword))

	mux.Get("/ping", http.HandlerFunc(ping))

	// serving static files
	fileServer := http.FileServer(http.Dir(app.staticFiles))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	tImgFileServer := http.FileServer(http.Dir(app.tournamentImages))
	mux.Get("/tournament/img/", http.StripPrefix("/tournament/img", tImgFileServer))

	uImgFileServer := http.FileServer(http.Dir(app.userImages))
	mux.Get("/user/img/", http.StripPrefix("/user/img", uImgFileServer))

	return standardMiddleware.Then(mux)
}
