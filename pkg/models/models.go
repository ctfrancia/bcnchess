package models

import (
	"errors"
	"time"
)

// ErrNoRecord holds a variable to return incase no record is found
var (
	ErrNoRecord           = errors.New("models: no matching record")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrEmailAlreadyExists = "Address is Already in use"
)

// Tournament defines how are records are being saved and received from the DB
type Tournament struct {
	ID                    int
	Title                 string
	Location              string
	TournamentDate        time.Time
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

// User defines how our user table is structured
type User struct {
	ID               int
	FirstName        string
	LastName         string
	Email            string
	Password         []byte
	Club             string
	EloStandard      string
	EloRapid         string
	LichessUsername  string
	ChesscomUsername string
	Created          time.Time
	Active           bool
}
