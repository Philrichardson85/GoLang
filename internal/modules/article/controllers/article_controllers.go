package controllers

import (
	ArticleService "demoBlog/internal/modules/article/services"
	"demoBlog/pkg/html"
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
	html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the is"})
	return
}

article, err := controller.articleService.Find(id)
if err != nil {
	html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
	return
}

html.Render(c, http.StatusOK, "modules/articles/html/show", gin.H{"title": "Show article", "article": article})

}

func (controller *Controller) Create( c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Hello world .."})
}