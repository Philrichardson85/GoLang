package migration

import (
	articleModel "demoBlog/internal/modules/article/models"
	userModel "demoBlog/internal/modules/user/models"
	"demoBlog/pkg/database"
	"fmt"
	"log"
)

func Migrate()  {
	db := database.Connection()
	err := db.AutoMigrate(&userModel.User{}, &articleModel.Article{})

	if err != nil {
		log.Fatal("can't migrate")
		return
	}

	fmt.Println("migration done...")
}