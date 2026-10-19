package repositories

import (
	ArticleModel "xzl_blog/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
	Create(article ArticleModel.Article) ArticleModel.Article
	Search(title string) ArticleModel.Article
	Delete(id int) error
	Update(article ArticleModel.Article) ArticleModel.Article
}
