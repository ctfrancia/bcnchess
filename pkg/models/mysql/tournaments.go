package mysql

import (
	"database/sql"
	"errors"

	"github.com/ctfrancia/bcnchess/pkg/models"
)

// TournamentModel allows sets up our method
type TournamentModel struct {
	DB *sql.DB
}

// Insert is used for inserting into our Tournament Table
func (m *TournamentModel) Insert(t models.Tournament) (int, error) {
	stmt := `INSERT INTO tournaments (
		title, location, matchTimeStart, matchTimeEnd, additionalInformation, isOnline, timeControl, tournamentType, rated, poster, created, expires
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, t.Title, t.Location, t.MatchTimeStart, t.MatchTimeEnd, t.AdditionalInformation, t.IsOnline, t.TimeControl, t.TournamentType, t.Rated, t.Poster, t.Expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Get takes the id of the tournament as an argument and returns tournament data
func (m *TournamentModel) Get(id int) (*models.Tournament, error) {
	/*
		stmt := `SELECT * FROM tournaments
			WHERE expires > UTC_TIMESTAMP() AND id = ?`
	*/
	stmt := `SELECT * FROM tournaments WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	t := &models.Tournament{}

	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Location,
		&t.MatchTimeStart,
		&t.MatchTimeEnd,
		&t.AdditionalInformation,
		&t.IsOnline,
		&t.TimeControl,
		&t.TournamentType,
		&t.Rated,
		&t.Poster,
		&t.Created,
		&t.Expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return t, nil
}

// Latest returns the first n newest tournaments
func (m *TournamentModel) Latest() ([]*models.Tournament, error) {
	stmt := `SELECT * FROM tournaments
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tournaments := []*models.Tournament{}
	for rows.Next() {
		t := &models.Tournament{}
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Location,
			&t.MatchTimeStart,
			&t.MatchTimeEnd,
			&t.AdditionalInformation,
			&t.IsOnline,
			&t.TimeControl,
			&t.TournamentType,
			&t.Rated,
			&t.Poster,
			&t.Created,
			&t.Expires,
		)
		if err != nil {
			return nil, err
		}
		tournaments = append(tournaments, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tournaments, nil
}
