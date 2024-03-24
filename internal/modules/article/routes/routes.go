package routes

import (
	articleCtrl "demoBlog/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

//set routes
func Routes(router *gin.Engine)  {
	articlesController := articleCtrl.New()
		router.GET("/articles/:id", articlesController.Show)
}