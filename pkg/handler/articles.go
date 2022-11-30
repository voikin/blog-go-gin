package handler

import (
	"log"
	"net/http"
	"strconv"
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

	ctx.AbortWithStatus(http.StatusOK)
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
