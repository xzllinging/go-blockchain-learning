package routes

import (
	"xzl_blog/internal/middlewares"
	articleCtrl "xzl_blog/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := articleCtrl.New()
	router.GET("/articles/:id", articlesController.Show)
	router.GET("/Comment", articlesController.Comment)
	router.GET("/Search", articlesController.Search)

	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", articlesController.Create)
		authGroup.POST("/store", articlesController.Store)
		authGroup.DELETE("/:id", articlesController.Delete)
		authGroup.GET("/:id/edit", articlesController.Edit)
		authGroup.POST("/:id/update", articlesController.Update)
	}
}
