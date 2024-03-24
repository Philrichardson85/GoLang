package controllers

import (
	ArticleService "demoBlog/internal/modules/article/services"
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

func (controller *Controller) Show( c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello article"})
}