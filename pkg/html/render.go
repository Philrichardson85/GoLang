package html

import (
	"demoBlog/internal/providers/view"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, html_code int, template_folder string, data gin.H)  {
	data = view.WithGlobalData(c, data)
	c.HTML(html_code, template_folder, data)
}