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
	SaveUser(user *models.User) (int64, error)
	DeleteUser(id int64) error
	GetUserByID(id int64) (*models.User, error)
	GetUserByNickname(nickname string) (*models.User, error)
	GetUsers() ([]*models.User, error)
}

type ArticleInfoRepository interface {
	SaveArticleInfo(article *models.ArticleInfo) (int64, error)
	DeleteArticleInfo(id int64) error
	GetArticleInfo(id int64) (*models.ArticleInfo, error)
}

type ArticleTextRepository interface {
	SaveArticleText(article *models.ArticleText) error
	DeleteArticleText(id int64) error
	GetArticlesText() ([]*models.ArticleText, error)
	GetArticleTextByID(id int64) (*models.ArticleText, error)
}

type SessionRepository interface {
	SaveSession(session *models.Session) (int64, error)
	DeleteSession(id int64) error
	GetSessionByToken(sessionToken string) (*models.Session, error)
	GetSessionByID(id int64) (*models.Session, error)
}

func NewRepository(db *sql.DB, client *elasticsearch.Client) *Repository {
	return &Repository{
		UserRepository:        mysql.NewUsersMySQL(db),
		ArticleInfoRepository: mysql.NewArticlesInfoMySQL(db),
		ArticleTextRepository: elastic_search.NewArticlesTextElasticSearch(client),
		SessionRepository:     mysql.NewSessionsMySQL(db),
	}
}
