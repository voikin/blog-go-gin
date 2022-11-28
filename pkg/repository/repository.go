package repository

import (
	"database/sql"
	"github.com/dazai404/blog-go-gin/models"
	"github.com/dazai404/blog-go-gin/pkg/elastic_search"
	"github.com/dazai404/blog-go-gin/pkg/mysql"
	"github.com/elastic/go-elasticsearch/v8"
)

type Repository struct {
	UserRepository
	ArticleInfoRepository
	ArticleTextRepository
	SessionRepository
}

type UserRepository interface {
	SaveUser(user *models.User) error
	DeleteUser(id int64) error
	GetUser(id int64) (*models.User, error)
}

type ArticleInfoRepository interface {
	SaveArticleInfo(article *models.ArticleInfo) error
	DeleteArticleInfo(id int64) error
	GetArticleInfo(id int64) (*models.ArticleInfo, error)
}

type ArticleTextRepository interface {
	SaveArticleText(article *models.ArticleText) error
	DeleteArticleText(id int64) error
	GetArticleText(id int64) (*models.ArticleText, error)
}

type SessionRepository interface {
	SaveSession(session *models.Session) error
	DeleteSession(id int64) error
	GetSession(id int64) (*models.Session, error)
}

func NewRepository(db *sql.DB, client *elasticsearch.Client) *Repository {
	return &Repository{
		UserRepository:        mysql.NewUsersMySQL(db),
		ArticleInfoRepository: mysql.NewArticlesInfoMySQL(db),
		ArticleTextRepository: elastic_search.NewArticlesTextElasticSearch(client),
		SessionRepository:     mysql.NewSessionsMySQL(db),
	}
}
