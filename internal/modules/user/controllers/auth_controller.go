package controllers

import (
	"demoBlog/internal/modules/user/requests/auth"
	UserService "demoBlog/internal/modules/user/services"
	"demoBlog/pkg/html"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller{
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register( c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register" ,gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister( c *gin.Context) {
	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.userService.Create(request)

	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	log.Printf("the user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}