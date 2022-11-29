package mysql

import (
	"database/sql"
	"fmt"
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

func (m *UsersMySQL) SaveUser(user *models.User) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (nickname, email, role, password_hash) VALUES (?, ?, ?, ?)", usersTable)

	row, err := m.db.Exec(query, user.Nickname, user.Email, user.Role, user.PasswordHash)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *UsersMySQL) DeleteUser(id int64) error {
	return nil
}

func (m *UsersMySQL) GetUser(id int64) (*models.User, error) {
	return nil, nil
}
