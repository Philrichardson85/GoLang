package controllers

import (
	"demoBlog/internal/modules/user/requests/auth"
	UserService "demoBlog/internal/modules/user/services"
	"demoBlog/pkg/old"
	"demoBlog/pkg/sessions"
	"demoBlog/pkg/converters"
	"demoBlog/pkg/errors"
	"demoBlog/pkg/html"
	"log"
	"net/http"
	"strconv"

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
		errors.Init()
		errors.SetFromErrors(err)

		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		
		old.Init()
		old.Set(c)
		sessions.Set(c, "old",converters.UrlValuesToString(old.Get()))


		c.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Email Address already exists")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		
		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.userService.Create(request)

	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("the user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)

		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		
		old.Init()
		old.Set(c)
		sessions.Set(c, "old",converters.UrlValuesToString(old.Get()))


		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userService.HandleUserLogin(request)

	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())

		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		
		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))


		c.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("the user logged in successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")
	c.Redirect(http.StatusFound, "/")
}