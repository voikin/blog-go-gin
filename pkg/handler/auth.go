package handler

import (
	"github.com/dazai404/blog-go-gin/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	input := &struct {
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		Password string `json:"password"`
	}{}

	err := ctx.BindJSON(input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user := &models.User{
		Nickname:     input.Nickname,
		Email:        input.Email,
		Role:         input.Role,
		PasswordHash: hash,
	}

	id, err := h.repository.SaveUser(user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
