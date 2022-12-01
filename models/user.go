package models

import "time"

type User struct {
	ID           int64     `json:"id"`
	Nickname     string    `json:"nickname" binding:"required"`
	Email        string    `json:"email" binding:"required, email"`
	Role         string    `json:"role" binding:"required"`
	PasswordHash []byte    `json:"passwordHash" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
}
