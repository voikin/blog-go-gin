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
		auth.POST("/sign-in", h.signIn)
	}

	admin := router.Group("/admin", h.adminMiddleware)
	{
		users := admin.Group("/users")
		{
			users.GET("/", h.getUsers)
			users.GET("/:id", h.getUserByID)
			users.DELETE("/:id", h.deleteUserByID)
		}
	}

	show := router.Group("/show", h.authMiddleware)
	{
		show.GET("/")
		show.GET("/:id")
	}

	api := router.Group("/api")
	{
		articles := api.Group("/articles", h.authMiddleware)
		{
			articles.POST("/new", h.saveArticle)
			articles.GET("/")
			articles.GET("/:id", h.getArticleByID)
			articles.DELETE("/:id")
		}
	}

	return router
}
