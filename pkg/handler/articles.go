package handler

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/dazai404/blog-go-gin/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) saveArticle(ctx *gin.Context) {
	user, ok := ctx.Keys["user"].(*models.User)
	if !ok {
		log.Println("error user")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	input := &struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}{}

	err := ctx.BindJSON(input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	articleInfo := &models.ArticleInfo{
		UserID:    user.ID,
		Title:     input.Title,
		CreatedAt: time.Now(),
	}

	id, err := h.repository.SaveArticleInfo(articleInfo)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	articleText := &models.ArticleText{
		ID:   id,
		Text: input.Text,
	}

	err = h.repository.SaveArticleText(articleText)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getArticleByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	articleInfo, err := h.repository.GetArticleInfo(int64(id))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	articleText, err := h.repository.GetArticleTextByID(int64(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	article := &struct {
		ID        int64     `json:"id"`
		UserID    int64     `json:"user_id"`
		Title     string    `json:"title"`
		Text      string    `json:"text"`
		CreatedAt time.Time `json:"created_at"`
	}{
		ID: int64(id),
		UserID: articleInfo.UserID,
		Title: articleInfo.Title,
		Text: articleText.Text,
		CreatedAt: articleInfo.CreatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func (h *Handler) getAllArticles(ctx *gin.Context) {
	articlesInfo, err := h.repository.GetArticlesInfo()
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}

	articles := make([]*models.Article, 0, 1)

	wg := sync.WaitGroup{}
	artMux := sync.Mutex{}
	errMux := sync.Mutex{}
	var errGor *error

	for _, val := range articlesInfo {
		wg.Add(1)
		go func (context *gin.Context, info *models.ArticleInfo, errGo *error){
			articleID := info.ID
			articleText, err := h.repository.GetArticleTextByID(articleID)
			if err != nil {
				errMux.Lock()
				errGo = &err
				errMux.Unlock()
			}
			artMux.Lock()
			articles = append(articles, &models.Article{
				ID: info.ID,
				UserID: info.UserID,
				Title: info.Title,
				Text: articleText.Text,
				CreatedAt: info.CreatedAt,
			})
			artMux.Unlock()
			wg.Done()
		}(ctx, val, errGor)
	}

	if errGor != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"err": errGor,
	})

}
