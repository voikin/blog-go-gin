package mysql

import (
	"database/sql"
	"github.com/dazai404/blog-go-gin/models"
)

type UsersMySQL struct {
	db *sql.DB
}

func NewUsersMySQL(db *sql.DB) *UsersMySQL {
	return &UsersMySQL{
		db: db,
	}
}

func (m *UsersMySQL) SaveUser(*models.User) error {
	return nil
}

func (m *UsersMySQL) DeleteUser(id int64) error {
	return nil
}

func (m *UsersMySQL) GetUser(id int64) (*models.User, error) {
	return nil, nil
}
