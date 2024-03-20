package routing

import (
	"demoBlog/internal/providers/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	//setup routing
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes()  {
	routes.RegisterRoutes(GetRouter())
}