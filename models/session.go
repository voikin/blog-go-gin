package models

import (
	"fmt"
	"time"
)

type Session struct {
	ID        int64     `json:"id"`
	Session   string    `json:"session"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Session) IsExpired() error {
	if time.Now().Add(time.Hour * 3).Local().After(s.CreatedAt.Local().Add(time.Hour * 24)) {
		return fmt.Errorf("session in expired on %v", time.Now().Add(3*time.Hour).Local().Sub(s.CreatedAt.Add(time.Minute)))
	}
	return nil
}
