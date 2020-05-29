package models

import (
	"errors"
	"time"
)

// ErrNoRecord holds a variable to return incase no record is found
var ErrNoRecord = errors.New("models: no matching record")

// Tournament defines how are records are being saved and received from the DB
type Tournament struct {
	ID                    int
	Title                 string
	Location              string
	MatchTimeStart        time.Time
	MatchTimeEnd          time.Time
	AdditionalInformation string
	IsOnline              bool
	TimeControl           string
	TournamentType        string
	Rated                 bool
	Poster                string
	Created               time.Time
	Expires               time.Time
}
