package main

import (
	"net/http"
	"testing"
)

func TestApiGetLatestTournaments(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	t.Run("successful fetch", func(t *testing.T) {
		status, header, body := ts.get(t, "/api/tournament/latest")

		if status != http.StatusOK {
			t.Errorf("want: %d; got: %d", http.StatusOK, status)
		}

		if header.Get("Content-Type") != "application/json" {
			t.Errorf("want: %s; got: %s", "application/json", header.Get("Content-Type"))
		}

		if body == nil {
			t.Errorf("Body should not be nil")
		}
	})
}
