package repositories

import (
	ArticleModel "demoBlog/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
}