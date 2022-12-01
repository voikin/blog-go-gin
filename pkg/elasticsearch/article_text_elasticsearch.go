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
		Index:      articlesTextIndex,
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
		Index: articlesTextIndex,
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
	response := &models.ElasticSearchGetResponse{}
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
	req := esapi.SearchRequest{
		Index: []string{
			articlesTextIndex,
		},
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
	res := &models.ElasticSearchSearchResponse{}
	err = json.Unmarshal([]byte(resJSON.String()), res)
	if err != nil {
		return nil, err
	}
	articlesText := make([]*models.ArticleText, 0, 1)
	for _, hit := range res.Hits.Hits {
		id, err := strconv.Atoi(hit.ID)
		if err != nil {
			return nil, err
		}
		articleTextTemp := &models.ArticleText{
			ID: int64(id),
			Text: hit.Source.Text,
		}
		articlesText = append(articlesText, articleTextTemp)
	}
	return articlesText, nil
}

func (es *ArticlesTextElasticSearch) GetArticlesTextByIDs(ids []int64) ([]*models.ArticleText, error) {
	idss := &struct {
		Values []int64 `json:"values"`
	}{
		ids,
	}
	query := &struct {
		IDs *struct{
			Values []int64 `json:"values"`
		} `json:"ids"`
	}{
		idss,
	}
	requestBody := *&models.ElasticSearchSearchPostRequest{
		Query: query,
	}
	bodyJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req := esapi.SearchRequest{
		Index: []string{
			articlesTextIndex,
		},
		Body: bytes.NewReader(bodyJson),
	}
	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return nil, err
	}

	defer func() {
		if res != nil && res.Body != nil {
			if err := res.Body.Close(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.String())
	}

	response := &models.ElasticSearchSearchResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	articlesText := make([]*models.ArticleText, 0, 1)
	fmt.Println(response.Hits.Hits[0])
	for _, hit := range response.Hits.Hits {
		id, err := strconv.Atoi(hit.ID)
		if err != nil {
			return nil, err
		}
		articleTextTemp := &models.ArticleText{
			ID: int64(id),
			Text: hit.Source.Text,
		}
		articlesText = append(articlesText, articleTextTemp)
	}
	return articlesText, nil
}

