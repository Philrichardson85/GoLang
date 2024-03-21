package routes

import (
	homeCtrl "demoBlog/internal/modules/home/controllers"

	"github.com/gin-gonic/gin"
)

//set routes
func Routes(router *gin.Engine)  {
	homeController := homeCtrl.New()
		router.GET("/", homeController.Index)
}