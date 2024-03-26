package routes

import (
	userCtrl "demoBlog/internal/modules/user/controllers"
	"demoBlog/internal/middlewares"

	"github.com/gin-gonic/gin"
)

//set routes
func Routes(router *gin.Engine)  {
	userController := userCtrl.New()

	guestGroup := router.Group("/")

	guestGroup.Use(middlewares.IsGuest()) 
	{
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")

	authGroup.Use(middlewares.IsAuth()) 
	{
		authGroup.POST("/logout", userController.HandleLogout)
	}	
}
