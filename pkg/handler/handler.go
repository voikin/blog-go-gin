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

	api := router.Group("/api", h.authMiddleware)
	{
		private := api.Group("/private", h.adminMiddleware)
		{
			users := private.Group("/users")
			{
				users.GET("/", h.getUsers)
				users.GET("/:id", h.getUserByID)
				users.DELETE("/:id", h.deleteUserByID)
			}
			privateArticles := private.Group("/articles")
			{
				privateArticles.GET("/", h.getAllArticlesTest)
			}
		}
		public := api.Group("/public")
		{
			user := public.Group("/user")
			{
				user.GET("/", h.getUser)
				user.DELETE("/", h.deleteUser)
				user.PUT("/", h.updateUser)
			}
			articles := public.Group("/articles")
			{
				articles.GET("/user", h.getUserArticles)
				articles.POST("/new", h.saveArticle)
			}
		}
	}

	return router
}
