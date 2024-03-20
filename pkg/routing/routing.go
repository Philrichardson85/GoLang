package routing

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	//setup routing
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}