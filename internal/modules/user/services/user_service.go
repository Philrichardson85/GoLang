package services

import (
	userModel "demoBlog/internal/modules/user/models"
	UserRepository "demoBlog/internal/modules/user/repositories"
	"demoBlog/internal/modules/user/requests/auth"
	UserResponse "demoBlog/internal/modules/user/responses"
	"errors"

	"golang.org/x/crypto/bcrypt"
)
type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest ) (UserResponse.User, error) {

	var response UserResponse.User
	var user userModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("Error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email 
	user.Password = string(hashedPassword)


	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating the user")
	}

	return UserResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExists(email string) bool{
	user := userService.userRepository.FindByEmail(email)

	if user.ID != 0 {
		return true
	}

	return false
}
