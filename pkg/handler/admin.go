package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUsers(ctx *gin.Context) {
	users, err := h.repository.GetUsers()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *Handler) getUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.repository.GetUserByID(int64(id))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *Handler) deleteUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.repository.DeleteUser(int64(id))

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.AbortWithStatus(http.StatusOK)
}
