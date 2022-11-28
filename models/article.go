package models

import "time"

type Article struct {
	ID          int64        `json:"id"`
	ArticleInfo *ArticleInfo `json:"article_info"`
	ArticleText *ArticleText `json:"article_text"`
}

type ArticleInfo struct {
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleText struct {
	Text string `json:"text"`
}
