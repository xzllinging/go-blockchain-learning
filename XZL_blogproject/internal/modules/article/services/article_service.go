package services

import (
	"errors"
	ArticleModel "xzl_blog/internal/modules/article/models"
	ArticleRepository "xzl_blog/internal/modules/article/repositories"
	"xzl_blog/internal/modules/article/requests/articles"
	ArticleResponse "xzl_blog/internal/modules/article/responses"
	UserResponse "xzl_blog/internal/modules/user/responses"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)

	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)

	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article

	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return ArticleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article
	var response ArticleResponse.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := articleService.articleRepository.Create(article)

	if newArticle.ID == 0 {
		return response, errors.New("error in creating the article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}

func (articleService *ArticleService) Search(title string) ArticleResponse.Article {
	articles := articleService.articleRepository.Search(title)

	return ArticleResponse.ToArticle(articles)
}

func (articleService *ArticleService) Delete(id int) error {
	// 先检查文章是否存在
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return errors.New("文章不存在")
	}

	// 执行删除操作
	return articleService.articleRepository.Delete(id)
}

func (articleService *ArticleService) UpdateAsUser(id int, request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article

	// 先检查文章是否存在
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("文章不存在")
	}

	// 检查是否是文章作者
	if article.UserID != user.ID {
		return response, errors.New("没有权限修改此文章")
	}

	// 更新文章内容
	article.Title = request.Title
	article.Content = request.Content

	updatedArticle := articleService.articleRepository.Update(article)
	if updatedArticle.ID == 0 {
		return response, errors.New("更新文章失败")
	}

	return ArticleResponse.ToArticle(updatedArticle), nil
}
