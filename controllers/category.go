package controllers

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
)

type CategoryController struct {
	BaseController
}

type DispCategory struct {
	models.Blog
	ChangeCategory bool
}

func (this *CategoryController) Index(c echo.Context) error {
	this.BeforeFilter(c)

	blog := models.NewBlog()
	category := models.NewCategory()
	categories := category.All()
	blog_data := blog.FindListOrderCategory()
	blogs := dispCategory(blog_data)
	setChangeCategoryFlg(blogs)

	this.SetResponse("BlogList", blogs)
	this.SetResponse("Categories", categories)

	this.MetaTitle = "TKO技術ブログ|カテゴリ別ブログ一覧"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "カテゴリ別ブログ一覧"
	this.MetaRobots = "noydir,noodp,index,follow"

	return this.Render(c, http.StatusOK, "category.html")
}

func dispCategory(blogs []*models.Blog) []*DispCategory {
	results := []*DispCategory{}

	for _, blog := range blogs {
		result := new(DispCategory)
		result.ID = blog.ID
		result.CreatedAt = blog.CreatedAt
		result.UpdatedAt = blog.UpdatedAt
		result.DeletedAt = blog.DeletedAt
		result.Title = blog.Title
		result.Body = blog.Body
		result.Keywords = blog.Keywords
		result.IsShow = blog.IsShow
		result.Category = blog.Category
		result.ReleaseDate = blog.ReleaseDate
		result.BlogCount = blog.BlogCount
		results = append(results, result)
	}

	return results
}

func setChangeCategoryFlg(blogs []*DispCategory) {
	beforeCategory := uint(0)
	for _, blog := range blogs {
		blog.ChangeCategory = beforeCategory != blog.Category

		beforeCategory = blog.Category
	}
}
