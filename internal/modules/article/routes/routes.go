package routes

import (
	articleCtrl "demoBlog/internal/modules/article/controllers"
	"demoBlog/internal/middlewares"
	"github.com/gin-gonic/gin"
)

//set routes
func Routes(router *gin.Engine)  {
	articlesController := articleCtrl.New()
		
	router.GET("/articles/:id", articlesController.Show)

	authGroup := router.Group("/articles")
		
	authGroup.Use(middlewares.IsAuth()) 
	{
		authGroup.GET("/create", articlesController.Create)
	}
}