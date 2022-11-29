package models

import "time"

type User struct {
	ID           int64     `json:"id"`
	Nickname     string    `json:"nickname"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	PasswordHash []byte    `json:"passwordHash"`
	CreatedAt    time.Time `json:"created_at"`
}
