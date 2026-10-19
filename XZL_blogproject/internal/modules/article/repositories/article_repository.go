package repositories

import (
	ArticleModel "xzl_blog/internal/modules/article/models"
	"xzl_blog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article

	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

func (articleRepository *ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}

func (articleRepository *ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article

	articleRepository.DB.Create(&article).Scan(&newArticle)

	return newArticle
}
func (articleRepository *ArticleRepository) Search(title string) ArticleModel.Article {
	var article ArticleModel.Article

	articleRepository.DB.Where("title LIKE ?", "%"+title+"%").First(&article)

	return article

}

func (articleRepository *ArticleRepository) Delete(id int) error {
	result := articleRepository.DB.Delete(&ArticleModel.Article{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleRepository *ArticleRepository) Update(article ArticleModel.Article) ArticleModel.Article {
	var updatedArticle ArticleModel.Article

	articleRepository.DB.Model(&article).Updates(ArticleModel.Article{
		Title:   article.Title,
		Content: article.Content,
	}).Scan(&updatedArticle)

	return updatedArticle
}
