package mock

import (
	"time"

	"github.com/ctfrancia/bcnchess/pkg/models"
)

var mockTournament = &models.Tournament{
	ID:                    1,
	Title:                 "Test tournament 1",
	Location:              "carrer alexandre gali Barcelona, Spain, 08000",
	TournamentDate:        time.Now(),
	MatchTimeStart:        time.Now(),
	MatchTimeEnd:          time.Now(),
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
type TournamentModel struct{}

// Insert mocks out Insert Tournament method
func (m *TournamentModel) Insert(t *models.Tournament) (int, error) {
	return 2, nil
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
