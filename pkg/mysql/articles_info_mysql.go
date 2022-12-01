package mysql

import (
	"database/sql"
	"fmt"
	"time"

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
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", articlesInfoTable)

	row := m.db.QueryRow(query, id)

	articleInfo := &models.ArticleInfo{}

	err := row.Scan(&articleInfo.ID, &articleInfo.UserID, &articleInfo.Title, &articleInfo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return articleInfo, nil
}

func (m *ArticlesInfoMySQL) GetArticlesInfo() ([]*models.ArticleInfo, error) {
	query := fmt.Sprintf("SELECT * FROM %s", articlesInfoTable)

	rows, err := m.db.Query(query)

	defer func() {
		if rows != nil {
			if err := rows.Close(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	if err != nil {
		return nil, err
	}

	articlesInfo := make([]*models.ArticleInfo, 0, 1)
	var id, userID int64
	var title string
	var createdAt time.Time

	for rows.Next() {
		err = rows.Scan(&id, &userID, &title, &createdAt)
		if err != nil {
			return nil, err
		}
		articlesInfo = append(articlesInfo, &models.ArticleInfo{
			ID: id,
			UserID: userID,
			Title: title,
			CreatedAt: createdAt,
		})
	}

	return articlesInfo, nil
}
