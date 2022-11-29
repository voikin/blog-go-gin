package mysql

import (
	"database/sql"
	"fmt"
	"github.com/dazai404/blog-go-gin/models"
	"time"
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
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", usersTable)

	row := m.db.QueryRow(query, id)

	return row.Err()
}

func (m *UsersMySQL) GetUserByNickname(nickname string) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE nickname = ?", usersTable)

	row := m.db.QueryRow(query, nickname)

	user := new(models.User)

	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Role, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (m *UsersMySQL) GetUserByID(id int64) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", usersTable)

	row := m.db.QueryRow(query, id)

	user := new(models.User)

	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Role, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UsersMySQL) GetUsers() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", usersTable)

	rows, err := m.db.Query(query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	users := make([]*models.User, 0, 1)
	var id int64
	var nickname, email, role string
	var passwordHash []byte
	var createdAt time.Time

	for rows.Next() {
		err = rows.Scan(&id, &nickname, &email, &role, &passwordHash, &createdAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &models.User{
			ID:           id,
			Nickname:     nickname,
			Email:        email,
			Role:         role,
			PasswordHash: passwordHash,
			CreatedAt:    createdAt,
		})
	}

	return users, nil
}
