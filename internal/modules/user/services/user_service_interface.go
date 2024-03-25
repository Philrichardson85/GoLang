package services

import (
	"demoBlog/internal/modules/user/requests/auth"
	UserResponse "demoBlog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
}