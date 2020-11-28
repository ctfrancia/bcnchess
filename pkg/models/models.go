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
	ErrEmailAlreadyExists = "Address is already in use"
)

// Tournament defines how are records are being saved and received from the DB
type Tournament struct {
	ID                    int
	Title                 string
	Location              string
	TournamentDate        time.Time
	MatchTimeStart        string
	MatchTimeEnd          string
	AdditionalInformation string
	IsOnline              bool
	TimeControl           string
	TournamentType        string
	Rated                 bool
	Poster                string
	TournamentContact     string
	Created               time.Time
	Expires               time.Time
}

// NewUserJSON is the JSON received in the request
type NewUserJSON struct {
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Club             string `json:"club"`
	ClubCountry      string `json:"clubCountry"`
	UserCountry      string `json:"userCountry"`
	LichessUsername  string `json:"lichessUsername"`
	ChesscomUsername string `json:"chesscomUsername"`
	// eloStandard      string `json:"eloStandard"`
	// eloRapid         string `json:"eloRapid"`
	// created          string `json:""`
	// active           bool   `json:""`
}

// User defines how our user table is structured
type User struct {
	ID               int
	FirstName        string
	LastName         string
	Email            string
	Password         []byte
	ClubName         string
	ClubCountry      string
	UserCountry      string
	EloStandard      string
	EloRapid         string
	LichessUsername  string
	ChesscomUsername string
	Created          time.Time
	Active           bool
}

// TournamentResponse defines our JSON response for a tournament
type TournamentResponse struct {
	ID                    int       `json:"id"`
	Title                 string    `json:"title"`
	Location              string    `json:"location"`
	TournamentDate        time.Time `json:"tournamentDate"`
	MatchTimeStart        string    `json:"matchTimeStart"`
	MatchTimeEnd          string    `json:"matchTimeEnd"`
	AdditionalInformation string    `json:"additionalInformation"`
	IsOnline              bool      `json:"isOnline"`
	TimeControl           string    `json:"timeControl"`
	TournamentType        string    `json:"tournamentType"`
	Rated                 bool      `json:"rated"`
	Poster                string    `json:"poster"`
	TournamentContact     string    `json:"tournamentContact"`
	Created               time.Time `json:"created"`
	Expires               time.Time `json:"expires"`
}

// UserResponse defines our JSON response for a user
type UserResponse struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	Club             string    `json:"club"`
	EloStandard      string    `json:"eloStandard"`
	EloRapid         string    `json:"eloRapid"`
	LichessUsername  string    `json:"lichessUsername"`
	ChesscomUsername string    `json:"chesscomUsername"`
	Created          time.Time `json:"created"`
	Active           bool      `json:"active"`
}
