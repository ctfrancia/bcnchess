package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"

	"github.com/ctfrancia/bcnchess/pkg/forms"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want: %d; got: %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestShowTournament(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/tournament/1", http.StatusOK, []byte("Additional information about the tournament here")},
		{"Non-existent ID", "/tournament/1000", http.StatusNotFound, nil},
		{"Negative ID", "/tournament/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/tournament/1.2", http.StatusNotFound, nil},
		{"String ID", "/tournament/foo", http.StatusNotFound, nil},
		{"Empty ID", "/tournament/", http.StatusNotFound, nil},
		{"Trailing Slash", "/tournament/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		code, _, body := ts.get(t, tt.urlPath)
		if code != tt.wantCode {
			t.Errorf("want: %d; got: %d", tt.wantCode, code)
		}

		if !bytes.Contains(body, tt.wantBody) {
			t.Errorf("body received %q", body)
		}
	}
}

func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	tests := []struct {
		name             string
		userFirstName    string
		userLastName     string
		userEmail        string
		userPassword     string
		csrfToken        string
		club             string
		eloStandard      string
		eloRapid         string
		lichessUsername  string
		chesscomUsername string
		wantCode         int
		wantBody         []byte
	}{
		{"Valid Submission", "John", "Doe", "john@example1.com", "validpa$$word", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, nil},
		{"Empty Name", "", "Doe", "john@example2.com", "validpa$$word", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(forms.ErrCannotBeBlank)},
		{"Empty Lastname", "John", "", "john@example3.com", "validpa$$word", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(forms.ErrCannotBeBlank)},
		{"Empty Email", "John", "Doe", "", "validpa$$word", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(forms.ErrCannotBeBlank)},
		{"Invalid Email(missing @)", "John", "Doe", "johnexample4.com", "validpa$$word", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(forms.ErrFieldInvalid)},
		{"Invalid Email(local)", "John", "Doe", "@example.com", "validpa$$word", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(forms.ErrFieldInvalid)},
		{"Invalid Email(incomplete address)", "John", "Doe", "joh@nexample.", "validpa$$word", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(forms.ErrFieldInvalid)},
		{"Short Password", "John", "Doe", "joh@nexample.com", "no", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(forms.ErrFieldTooShort(6))},
		// {"Duplicate Email", "John", "Doe", "dupe@example.com", "no", csrfToken, "Congres", "1500", "1300", "exampleLichess", "chesscomExample", http.StatusOK, []byte(models.ErrEmailAlreadyExists)},
		{"Invalid CSRF Token", "", "", "", "no", "wrongToken", "", "", "", "", "", http.StatusBadRequest, nil},
	}
	// TODO: Solve failing test that's commented above
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("firstName", tt.userFirstName)
			form.Add("lastName", tt.userLastName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)
			form.Add("club", tt.club)
			form.Add("eloStandard", tt.eloStandard)
			form.Add("eloRapid", tt.eloRapid)
			form.Add("lichessUsername", tt.lichessUsername)
			form.Add("chesscomUsername", tt.chesscomUsername)

			code, _, body := ts.postForm(t, "/user/signup", form)

			if code != tt.wantCode {
				t.Errorf("want: %d; got: %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body %s to contain %q", body, tt.wantBody)
			}
		})
	}
}

func TestCreateTournamentForm(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	t.Run("Unauthenticated", func(t *testing.T) {
		code, headers, _ := ts.get(t, "/tournament/create")
		if code != http.StatusSeeOther {
			t.Errorf("wanted: %d; got: %d", http.StatusSeeOther, code)
		}
		if headers.Get("Location") != "/user/login" {
			t.Errorf("wants: %s; got %s", "/user/login", headers.Get("Location"))
		}
	})

	t.Run("Authenticated", func(t *testing.T) {
		_, _, body := ts.get(t, "/user/login")
		csrfToken := extractCSRFToken(t, body)

		form := url.Values{}

		form.Add("email", "jdoe@example.com")
		form.Add("password", "")
		form.Add("csrf_token", csrfToken)
		ts.postForm(t, "/user/login", form)

		code, _, body := ts.get(t, "/tournament/create")
		if code != 200 {
			t.Errorf("want: %d; got: %d", 200, code)
		}

		formTag := "<form action='/tournament/create' enctype='multipart/form-data' method='POST'>"
		if !bytes.Contains(body, []byte(formTag)) {
			t.Errorf("want body %s to contain %s", body, formTag)
		}
	})
}

/*
func TestAddUserToTournament(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	t.Run("Successful Update", func(t *testing.T) {
		err := ts.put() // AddUserToTournament(1, 1)
		if err != nil {
			t.Error("error when adding foreign key in tournament table with a user id", err)
		}
	})
}
*/

/*
func TestGetTournament(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: Skipping integration test")
	}
	db, teardown := newTestDB(t)
	defer teardown()

	m := TournamentModel{db}

	_, err := m.Get(1)

	if err != nil {
		t.Error(err)
	}
}
*/
/*
func TestInsertTournament(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	if testing.Short() {
		t.Skip("mysql: Skipping integration test")
	}
	tests := []struct {
		name           string
		wantTournament *models.Tournament
		wantError      error
	}{
		{
			name: "Valid Insert",
			wantTournament: &models.Tournament{
				Title:                 "Title 1",
				Location:              "some location",
				TournamentDate:        time.Now(),
				MatchTimeStart:        "1400",
				MatchTimeEnd:          "never",
				AdditionalInformation: "additional information",
				IsOnline:              true,
				TimeControl:           "5+0",
				TournamentType:        "Swiss",
				Rated:                 false,
				Poster:                "none",
				// Expires:               time.Now().Add(time.Hour + 1),
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TournamentModel{db}
			_, err := m.Insert(tt.wantTournament)

			if err != nil {
				t.Fatal(err)
			}
		})
	}

}
*/
