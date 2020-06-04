package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/ctfrancia/bcnchess/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// UserModel defines the data associated with the UserModel
type UserModel struct {
	DB *sql.DB
}

// Insert takes the User model to insert into the user table
func (m *UserModel) Insert(u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (
		firstName,
		lastName,
		email,
		password,
		club,
		eloStandard,
		eloRapid,
		created
		) VALUES (?, ?, ?, ?, ?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, u.FirstName, u.LastName, u.Email, string(hashedPassword), u.Club, u.EloStandard, u.EloRapid)
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_us_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

// Authenticate takes the email and password to check if information is valid
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, password FROM users where email = ? AND active = TRUE"
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	return id, nil
}

// Get will fetch a single record of a user from the user table
func (m *UserModel) Get(id int) (*models.User, error) {
	u := &models.User{}

	stmt := "SELECT id, firstName, lastName, email, club, eloStandard, eloRapid, lichessUsername, chesscomUserName, created, active FROM users WHERE id = ?"
	err := m.DB.QueryRow(stmt, id).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Club,
		&u.EloStandard,
		&u.EloRapid,
		&u.LichessUsername,
		&u.ChesscomUsername,
		&u.Created,
		&u.Active,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return u, nil
}

// UpdatePassword updates the user's password
func (m *UserModel) UpdatePassword(id int, pw string) error {
	hp, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	if err != nil {
		return err
	}
	stmt := `UPDATE users SET password = ? WHERE id = ?`
	_, err = m.DB.Exec(stmt, string(hp), id)

	if err != nil {
		return err
	}
	return nil
}
