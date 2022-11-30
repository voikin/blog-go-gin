package elastic_search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dazai404/blog-go-gin/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strconv"
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
	jsonText, err := json.Marshal(&struct {
		Text string `json:"text"`
	}{articleText.Text})
	req := esapi.IndexRequest{
		Index:      "article_text",
		DocumentID: strconv.FormatInt(articleText.ID, 10),
		Body:       bytes.NewReader(jsonText),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), es.client)
	defer func() {
		if res != nil && res.Body != nil {
			if err := res.Body.Close(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
	if err != nil {
		return err
	}
	if res.IsError() {
		return errors.New(fmt.Sprintf("%s ERROR indexing document ID=%d error:%s", res.Status(), articleText.ID, res.String()))
	}
	return nil
}

func (es *ArticlesTextElasticSearch) DeleteArticleText(id int64) error {
	return nil
}
func (es *ArticlesTextElasticSearch) GetArticleText(id int64) (*models.ArticleText, error) {
	return nil, nil
}
