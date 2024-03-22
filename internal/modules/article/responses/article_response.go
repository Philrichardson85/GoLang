package responses

import (
	articleModel "demoBlog/internal/modules/article/models"
	userResponses "demoBlog/internal/modules/user/responses"
	"fmt"
)

type Article struct {
	ID uint
	Title string
	Content string
	Image string
	CreatedAt string
	User userResponses.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModel.Article) Article {
	return Article{
		ID: article.ID,
		Title: article.Title,
		Content: article.Content,
		Image: "/assets/img/demopic/10.jpg",
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
	User: userResponses.ToUser(article.User),
	}
}

func ToArticles(articles []articleModel.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}