package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ctfrancia/bcnchess/pkg/models"
	// "github.com/rs/cors"

	"github.com/justinas/nosurf"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XXS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

func (app *application) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			app.session.Put(r, "redirectPathAfterLogin", r.URL.Path)
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exists := app.session.Exists(r, "authenticatedUserID")

		if !exists {
			next.ServeHTTP(w, r)
			return
		}

		//TODO: implement the new authentication here.

		/*
			user, err := app.users.Get(app.session.GetInt(r, "authenticatedUserID"))
			if errors.Is(err, models.ErrNoRecord) || !user.Active {
				app.session.Remove(r, "authenticatedUserID")
				next.ServeHTTP(w, r)
				return
			} else if err != nil {
				app.serverError(w, err)
				return
			}
			ctx := context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
		*/

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

/*
func (app *application) apiConsumption(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    fmt.Println("api consumption", r)
    c := cors.New(cors.Options{
      AllowedOrigins: []string{"http://localhost:8080"},
	  })

    next.ServeHTTP(w, r)
  })
}
*/
