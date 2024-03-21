package seeder

import (
	articleModel "demoBlog/internal/modules/article/models"
	userModel "demoBlog/internal/modules/user/models"
	"demoBlog/pkg/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed()  {
	db := database.Connection()
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 12)

	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := userModel.User{Name: "Random Name", Email: "email@email.com", Password: string(hashedPassword)} 
	db.Create(&user) // pass pointer of data to Create

	log.Printf("User created successfully with email: %s \n", user.Email)

	for i := 1; i <= 10; i++{
		article := articleModel.Article{Title: fmt.Sprintf("Random Title %d", i), Content: fmt.Sprintf("Random Content %d", i), UserID: 1} 
		db.Create(&article) // pass pointer of data to Create

		log.Printf("Article created successfully with Title: %s \n", article.Title)
	}
}