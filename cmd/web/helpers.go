package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	if app.serverConfig.Debug {
		http.Error(w, trace, http.StatusInternalServerError)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.defaultData(td, r))
	if err != nil {
		app.serverError(w, err)
	}

	buf.WriteTo(w)
}

func (app *application) defaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.session.PopString(r, "flash")
	td.CurrentYear = time.Now().Year()
	td.IsAuthenticated = app.isAuthenticated(r)

	return td
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}
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
