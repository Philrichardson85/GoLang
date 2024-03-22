package repositories

import (
	ArticleModel "demoBlog/internal/modules/article/models"
	"demoBlog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository {
		DB: database.Connection(),
	}
}
func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}