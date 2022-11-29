package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) authMiddleware(ctx *gin.Context) {
	sessionToken, err := ctx.Cookie("session_cookie")
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	session, err := h.repository.GetSessionByToken(sessionToken)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if session.IsExpired() {
		ctx.SetCookie("session_cookie", "", 0, "/", "localhost", false, true)
		err = h.repository.DeleteSession(session.ID)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("session is expired"))
		return
	}

	user, err := h.repository.GetUserByID(session.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("session is expired"))
		return
	}

	ctx.Set("user", user)
}

func (h *Handler) adminMiddleware(ctx *gin.Context) {
	sessionToken, err := ctx.Cookie("session_cookie")
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	session, err := h.repository.GetSessionByToken(sessionToken)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if session.IsExpired() {
		ctx.SetCookie("session_cookie", "", 0, "/", "localhost", false, true)
		err = h.repository.DeleteSession(session.ID)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("session is expired"))
		return
	}

	user, err := h.repository.GetUserByID(session.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("session is expired"))
		return
	}

	if user.Role == "admin" {
		ctx.Set("user", user)
		ctx.Next()
		return
	}

	ctx.AbortWithError(http.StatusForbidden, errors.New("incorrect role"))
}
