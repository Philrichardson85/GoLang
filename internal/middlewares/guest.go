package middlewares

import (
	"github.com/gin-gonic/gin"
	"demoBlog/pkg/sessions"
	"net/http"
		"strconv"
		UserRepository "demoBlog/internal/modules/user/repositories"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Next()
	}
}