package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"xzl_blog/internal/modules/article/requests/articles"
	ArticleService "xzl_blog/internal/modules/article/services"
	"xzl_blog/internal/modules/user/helpers"
	"xzl_blog/pkg/converters"
	"xzl_blog/pkg/errors"
	"xzl_blog/pkg/html"
	"xzl_blog/pkg/old"
	"xzl_blog/pkg/sessions"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	// Get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// Find the article from the database
	article, err := controller.articleService.Find(id)

	// If the article is not found, show error page
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// if article found, render article template
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article", "article": article})
}

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{"title": "Create article"})
}
func (controller *Controller) Search(c *gin.Context) {
	title := c.Query("title")
	article := controller.articleService.Search(title)
	html.Render(c, http.StatusOK, "modules/article/html/search", gin.H{
		"title":   "Create article",
		"article": article,
	})
}
func (controller *Controller) Comment(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/comments", gin.H{"title": "Comment article"})

}

func (controller *Controller) Store(c *gin.Context) {
	// validate the request
	var request articles.StoreRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	// Create the article
	article, err := controller.articleService.StoreAsUser(request, user)

	// Check if there is any error on the article creation
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}

func (controller *Controller) Delete(c *gin.Context) {
	// 获取文章ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 删除文章
	err = controller.articleService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}

func (controller *Controller) Edit(c *gin.Context) {
	// 获取文章ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// 获取文章
	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/edit", gin.H{
		"title":   "编辑文章",
		"article": article,
	})
}

func (controller *Controller) Update(c *gin.Context) {
	// 获取文章ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 验证请求
	var request articles.StoreRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d/edit", id))
		return
	}

	user := helpers.Auth(c)

	// 更新文章
	article, err := controller.articleService.UpdateAsUser(id, request, user)
	if err != nil {
		c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d/edit", id))
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
