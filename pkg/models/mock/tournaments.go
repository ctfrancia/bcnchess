package mock

import (
	"database/sql"
	"time"

	"github.com/ctfrancia/bcnchess/pkg/models"
)

var mockTournament = &models.Tournament{
	ID:                    1,
	Title:                 "Test tournament 1",
	Location:              "carrer alexandre gali Barcelona, Spain, 08000",
	TournamentDate:        time.Now(),
	MatchTimeStart:        "1400",
	MatchTimeEnd:          "1500",
	AdditionalInformation: "Additional information about the tournament here",
	IsOnline:              false,
	TimeControl:           "5+0",
	TournamentType:        "Swiss",
	Rated:                 true,
	Poster:                "Path/To/Poster.png",
	Created:               time.Now(),
	Expires:               time.Now(),
}

// TournamentModel is used for testing our db methods
type TournamentModel struct {
	DB *sql.DB
}

// Insert mocks out Insert Tournament method
func (m *TournamentModel) Insert(t *models.Tournament) (int, error) {
	/*
		stmt := `INSERT INTO tournaments (
			title, location, tournamentDate, matchTimeStart, matchTimeEnd, additionalInformation, isOnline, timeControl, tournamentType, rated, poster, expires
			) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
		result, err := m.DB.Exec(stmt, t.Title, t.Location, t.TournamentDate, t.MatchTimeStart, t.MatchTimeEnd, t.AdditionalInformation, t.IsOnline, t.TimeControl, t.TournamentType, t.Rated, t.Poster, t.Expires)
		if err != nil {
			return 0, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return int(id), nil
	*/
	return 1, nil
}

// Get mocks out Get tournament method
func (m *TournamentModel) Get(id int) (*models.Tournament, error) {
	switch id {
	case 1:
		return mockTournament, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest mocks our Latest method for tournaments
func (m *TournamentModel) Latest() ([]*models.Tournament, error) {
	return []*models.Tournament{mockTournament}, nil
}
