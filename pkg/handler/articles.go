package handler

import (
	"github.com/dazai404/blog-go-gin/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
