package middlewares

import (
	"github.com/gin-gonic/gin"
	"demoBlog/pkg/sessions"
	"strconv"
	"net/http"
	UserRepository "demoBlog/internal/modules/user/repositories"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New()
	
	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Next()
	}
}