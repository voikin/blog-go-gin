package mysql

import (
	"database/sql"
	"fmt"
	"github.com/dazai404/blog-go-gin/models"
)

type SessionsMySQL struct {
	db *sql.DB
}

func NewSessionsMySQL(db *sql.DB) *SessionsMySQL {
	return &SessionsMySQL{
		db: db,
	}
}

func (m *SessionsMySQL) SaveSession(session *models.Session) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (session, user_id) VALUES (?, ?)", sessionsTable)

	res, err := m.db.Exec(query, session.Session, session.UserID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SessionsMySQL) DeleteSession(id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", sessionsTable)

	_, err := m.db.Exec(query, id)

	return err
}

func (m *SessionsMySQL) GetSessionByToken(sessionToken string) (*models.Session, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE session = ?", sessionsTable)

	row := m.db.QueryRow(query, sessionToken)

	session := new(models.Session)

	err := row.Scan(&session.ID, &session.Session, &session.UserID, &session.CreatedAt)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (m *SessionsMySQL) GetSessionByID(id int64) (*models.Session, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", sessionsTable)

	row := m.db.QueryRow(query, id)

	session := new(models.Session)

	err := row.Scan(&session.ID, &session.Session, &session.UserID, &session.CreatedAt)
	if err != nil {
		return nil, err
	}

	return session, nil
}
