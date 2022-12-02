package handler

import (
	"github.com/dazai404/blog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUserArticles(ctx *gin.Context) {
	user, ok := ctx.Keys["user"].(*models.User)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "incorrect user")
		return
	}
	userID := user.ID
	userArticlesInfo, err := h.repository.GetUserArticlesInfo(userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	articles := make([]*models.Article, 0, 1)
	ids := make([]int64, 0, 1)

	for _, val := range userArticlesInfo {
		ids = append(ids, val.ID)
	}

	articlesText, err := h.repository.GetArticlesTextByIDs(ids)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	for i, val := range userArticlesInfo {
		if userArticlesInfo[i].ID != articlesText[i].ID {
			newErrorResponse(ctx, http.StatusInternalServerError, "error with connecting entities from databases")
			return
		}
		articles = append(articles, &models.Article{
			ID:        val.ID,
			UserID:    val.UserID,
			Title:     val.Title,
			Text:      articlesText[i].Text,
			CreatedAt: val.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
