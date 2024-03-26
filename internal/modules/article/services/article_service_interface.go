package services

import (
	"demoBlog/internal/modules/article/requests/articles"
	ArticleResponse "demoBlog/internal/modules/article/responses"
	UserResponse "demoBlog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}