package routes

import (
	"demoBlog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

//set routes
func Routes(router *gin.Engine)  {
		router.GET("/", func(c *gin.Context) {
			html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
				"title": "Phil's Home website",
			})
		})
}