package mysql

import (
	"database/sql"
	"github.com/dazai404/blog-go-gin/models"
)

type ArticlesInfoMySQL struct {
	db *sql.DB
}

func NewArticlesInfoMySQL(db *sql.DB) *ArticlesInfoMySQL {
	return &ArticlesInfoMySQL{
		db: db,
	}
}

func (m *ArticlesInfoMySQL) SaveArticleInfo(articleInfo *models.ArticleInfo) error {
	return nil
}

func (m *ArticlesInfoMySQL) DeleteArticleInfo(id int64) error {
	return nil
}
func (m *ArticlesInfoMySQL) GetArticleInfo(id int64) (*models.ArticleInfo, error) {
	return nil, nil
}
