package handler

import (
	"github.com/dazai404/blog-go-gin/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		articles := api.Group("/articles")
		{
			articles.POST("/new")
			articles.GET("/:id")
			articles.DELETE("/:id")
		}
	}

	return router
}
