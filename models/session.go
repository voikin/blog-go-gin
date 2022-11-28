package models

import "time"

type Session struct {
	ID        int64     `json:"id"`
	Session   string    `json:"session"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
