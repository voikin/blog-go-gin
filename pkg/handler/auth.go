package handler

import (
	"github.com/dazai404/blog-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
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

func (h *Handler) signIn(ctx *gin.Context) {
	input := &struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}{}

	err := ctx.BindJSON(input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.repository.GetUserByNickname(input.Nickname)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(input.Password))
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	sessionToken := uuid.NewString()

	newSession := &models.Session{
		Session:   sessionToken,
		UserID:    user.ID,
		CreatedAt: time.Now(),
	}

	id, err := h.repository.SaveSession(newSession)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session, err := h.repository.GetSessionByID(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.SetCookie("session_cookie", sessionToken, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"user":    user,
		"session": session,
	})

}
