package controllers

import (
	"fmt"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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

	left_menu_categories := category.FindHasData()
	this.SetResponse("CategoryMenu", left_menu_categories)

	this.SetResponse("BlogList", blogs)
	this.SetResponse("Categories", categories)

	this.MetaTitle = "tko blogs|カテゴリ別ブログ一覧"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "カテゴリ別ブログ一覧"
	this.MetaRobots = "noydir,noodp,index,follow"

	return this.Render(c, http.StatusOK, "category.html")
}

func (this *CategoryController) One(c echo.Context) error {
	this.BeforeFilter(c)

	in_category, err := strconv.Atoi(c.Param("category_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	category := models.NewCategory()
	categories := category.All()
	blog_data := blog.FindListByCategory(uint(in_category))
	blogs := dispCategory(blog_data)
	setChangeCategoryFlg(blogs)

	left_menu_categories := category.FindHasData()
	this.SetResponse("CategoryMenu", left_menu_categories)

	this.SetResponse("BlogList", blogs)
	this.SetResponse("Categories", categories)

	this.MetaTitle = "tko blogs|カテゴリ別ブログ一覧"
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
