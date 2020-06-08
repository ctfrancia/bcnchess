package mock

import (
	"time"

	"github.com/ctfrancia/bcnchess/pkg/models"
)

var mockUser = &models.User{
	ID:               1,
	FirstName:        "Jane",
	LastName:         "Doe",
	Email:            "jdoe@example.com",
	Club:             "test",
	EloStandard:      "1000",
	EloRapid:         "1000",
	LichessUsername:  "test",
	ChesscomUsername: "test",
	Created:          time.Now(),
	Active:           true,
}

// UserModel is an empty interface so we can use methods against it
type UserModel struct{}

// Insert mocks our Insert method for User table
func (m *UserModel) Insert(u *models.User) error {
	switch u.Email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// Authenticate mocks out Authenticate method for user table
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "jdoe@example.com":
		return 1, nil

	default:
		return 0, models.ErrInvalidCredentials
	}
}

// Get mocks our Get method for user table
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil

	default:
		return nil, models.ErrNoRecord
	}
}

// UpdatePassword mocks out UpdatePassword for user table
func (m *UserModel) UpdatePassword(id int, pw string) error {
	return nil
}

// AddUserToTournament mocks our normal function
func (m *UserModel) AddUserToTournament(tID, uID int) error {
	return nil
}
