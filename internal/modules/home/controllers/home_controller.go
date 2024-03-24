package controllers

import (
	ArticleService "demoBlog/internal/modules/article/services"
	"demoBlog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller{
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller*Controller) Index( c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Phil's Home website",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories": controller.articleService.GetStoriesArticles(),
	})
}