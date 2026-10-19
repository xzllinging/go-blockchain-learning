package services

import (
	"xzl_blog/internal/modules/article/requests/articles"
	ArticleResponse "xzl_blog/internal/modules/article/responses"
	UserResponse "xzl_blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
	Search(title string) ArticleResponse.Article
	Delete(id int) error
	UpdateAsUser(id int, request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
