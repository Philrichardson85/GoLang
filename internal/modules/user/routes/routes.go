package routes

import (
	userCtrl "demoBlog/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

//set routes
func Routes(router *gin.Engine)  {
	userController := userCtrl.New()
		router.GET("/register", userController.Register)
		router.POST("/register", userController.HandleRegister)

		router.GET("/login", userController.Login)
		router.POST("/login", userController.HandleLogin)

		router.POST("/logout", userController.HandleLogout)
}
