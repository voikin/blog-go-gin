package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersTable        = "users"
	articlesInfoTable = "articles"
	sessionsTable     = "sessions"
)

func NewMySQLConnection(username string, password string, port string, dbname string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?parseTime=true", username, password, port, dbname))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
