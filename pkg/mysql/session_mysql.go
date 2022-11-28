package mysql

import (
	"database/sql"
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

func (m *SessionsMySQL) SaveSession(session *models.Session) error {
	return nil
}

func (m *SessionsMySQL) DeleteSession(id int64) error {
	return nil
}
func (m *SessionsMySQL) GetSession(id int64) (*models.Session, error) {
	return nil, nil
}
