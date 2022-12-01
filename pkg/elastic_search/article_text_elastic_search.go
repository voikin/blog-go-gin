package elastic_search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dazai404/blog-go-gin/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
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
	}{
		articleText.Text,
	})
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index:      "articles_text",
		DocumentID: strconv.FormatInt(articleText.ID, 10),
		Body:       bytes.NewReader(jsonText),
	}
	res, err := req.Do(context.Background(), es.client)
	fmt.Println(res.String())
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
		return fmt.Errorf("%s ERROR indexing document ID=%d error:%s", res.Status(), articleText.ID, res.String())
	}
	return nil
}

func (es *ArticlesTextElasticSearch) DeleteArticleText(id int64) error {
	return nil
}

func (es *ArticlesTextElasticSearch) GetArticleTextByID(id int64) (*models.ArticleText, error) {
	req := esapi.GetRequest{
		Index: "articles_text",
		DocumentID: strconv.FormatInt(id, 10),
		Pretty: true,
	}
	res, err := req.Do(context.Background(), es.client)
	fmt.Println(res.String())
	defer func() {
		if res != nil && res.Body != nil {
			if err := res.Body.Close(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
	if err != nil {
		log.Println(err.Error(), 1)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.String())
	}
	response := &models.ElasticSearchGet{}
	split := strings.SplitAfterN(res.String(), "] ", 2)
	body := split[1]
	fmt.Println(body)
	err = json.Unmarshal([]byte(body), response)
	if err != nil {
		log.Println(err.Error(), 2)
		return nil, err
	}
	articleText := &models.ArticleText{
		ID: id,
		Text: response.Source.Text,
	}
	return articleText, nil
}

func (es *ArticlesTextElasticSearch) GetArticlesText() ([]*models.ArticleText, error) {
	body := make(map[string]interface{})
	jsonMessage, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req := esapi.SearchRequest{
		Index: []string{
			"articles_text",
		},
		Body: bytes.NewReader(jsonMessage),
	}
	resJSON, err := req.Do(context.Background(), es.client)
	defer func() {
		if resJSON != nil && resJSON.Body != nil {
			if err := resJSON.Body.Close(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
	if err != nil {
		return nil, err
	}
	if resJSON.IsError() {
		return nil, errors.New(resJSON.String())
	}
	res := &models.ElasticSearchSearch{}
	err = json.Unmarshal([]byte(resJSON.String()), res)
	if err != nil {
		return nil, err
	}
	articlesText := make([]*models.ArticleText, 0, 1)
	for _, hit := range res.Hits.Hits {
		articleTextTemp := &models.ArticleText{
			ID: hit.ID,
			Text: hit.Source.Text,
		}
		articlesText = append(articlesText, articleTextTemp)
	}
	return articlesText, nil
}
