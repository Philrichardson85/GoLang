package controllers

import (
	ArticleService "demoBlog/internal/modules/article/services"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))
if err != nil {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the is"})
	return
}

article, err := controller.articleService.Find(id)
if err != nil {
	c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	return
}

	c.JSON(http.StatusOK, gin.H{"article": article})
}