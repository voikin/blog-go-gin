package models

import "time"

type User struct {
	ID           int64
	Nickname     string
	Email        string
	Role         string
	PasswordHash []byte
	CreatedAt    time.Time
}
