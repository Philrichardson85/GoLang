package controllers

import (
	ArticleRepository "demoBlog/internal/modules/article/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *Controller{
	return &Controller{
		articleRepository: ArticleRepository.New(),
	}
}

func (controller*Controller) Index( c *gin.Context) {
	// html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	// 	"title": "Phil's Home website",
	// })

	c.JSON(http.StatusOK, gin.H{
		"articles": controller.articleRepository.List(8)})
}