package elastic_search

import (
	"github.com/dazai404/blog-go-gin/models"
	"github.com/elastic/go-elasticsearch/v8"
)

type ArticlesTextElasticSearch struct {
	client *elasticsearch.Client
}

func NewArticlesTextElasticSearch(client *elasticsearch.Client) *ArticlesTextElasticSearch {
	return &ArticlesTextElasticSearch{
		client: client,
	}
}

func (es *ArticlesTextElasticSearch) SaveArticleText(articleText *models.ArticleText) error {
	return nil
}

func (es *ArticlesTextElasticSearch) DeleteArticleText(id int64) error {
	return nil
}
func (es *ArticlesTextElasticSearch) GetArticleText(id int64) (*models.ArticleText, error) {
	return nil, nil
}
