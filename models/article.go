package models

import "time"

type Article struct {
	ArticleInfo *ArticleInfo `json:"article_info"`
	ArticleText *ArticleText `json:"article_text"`
}

type ArticleInfo struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleText struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}
