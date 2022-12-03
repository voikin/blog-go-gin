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

func (h *Handler) getUser(ctx *gin.Context) {
	user, ok := ctx.Keys["user"].(*models.User)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	user, ok := ctx.Keys["user"].(*models.User)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user not found")
		return
	}
	err := h.repository.DeleteUser(user.ID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}

func (h *Handler) updateUser(ctx *gin.Context) {
	user, ok := ctx.Keys["user"].(*models.User)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user not found")
		return
	}
	input := &struct {
		Nickname *string `json:"nickname"`
		Email    *string `json:"email"`
	}{}

	err := ctx.BindJSON(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if input.Nickname != nil {
		user.Nickname = *input.Nickname
	}
	if input.Nickname != nil {
		user.Nickname = *input.Nickname
	}

	err = h.repository.UpdateUser(user)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
