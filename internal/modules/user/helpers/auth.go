package helpers

import (
	UserResponse "demoBlog/internal/modules/user/responses"
	UserRepository "demoBlog/internal/modules/user/repositories"
	"github.com/gin-gonic/gin"
	"demoBlog/pkg/sessions"
	"strconv"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User
	
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	var userRepo = UserRepository.New()
	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return response
	}

	return UserResponse.ToUser(user)
}