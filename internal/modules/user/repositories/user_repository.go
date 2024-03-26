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

func (userRepository *UserRepository) FindByEmail(email string) UserModel.User {
	var user UserModel.User

	userRepository.DB.First(&user, "email = ?", email)
	
	return user
}

func (userRepository *UserRepository) FindByID(id int) UserModel.User {
	var user UserModel.User

	userRepository.DB.First(&user, "id = ?", id)

	return user
}
