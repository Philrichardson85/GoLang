package services

import (
	ArticleModel "demoBlog/internal/modules/article/models"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() []ArticleModel.Article
	GetStoriesArticles() []ArticleModel.Article
}