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
	return 0, nil
}

// Get will fetch a single record of a user from the user table
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
