package mysql

import (
	"database/sql"
	"fmt"
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

func (m *ArticlesInfoMySQL) SaveArticleInfo(articleInfo *models.ArticleInfo) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, title) VALUES (?, ?)", articlesInfoTable)

	row, err := m.db.Exec(query, articleInfo.UserID, articleInfo.Title)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *ArticlesInfoMySQL) DeleteArticleInfo(id int64) error {
	return nil
}
func (m *ArticlesInfoMySQL) GetArticleInfo(id int64) (*models.ArticleInfo, error) {
	return nil, nil
}
