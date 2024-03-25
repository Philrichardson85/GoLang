package repositories

import (
	UserModel "demoBlog/internal/modules/user/models"
	"demoBlog/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository {
		DB: database.Connection(),
	}
}
func (userRepository *UserRepository) Create(user UserModel.User) UserModel.User {
	var newUser UserModel.User

	userRepository.DB.Create(&user).Scan(&newUser)
	
	return newUser
}