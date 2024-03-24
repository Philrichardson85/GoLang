package routes

import (
	articleRoutes "demoBlog/internal/modules/article/routes"
	homeRoutes "demoBlog/internal/modules/home/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
}
