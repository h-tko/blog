package controllers

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Index(c echo.Context) error {
	this.BeforeFilter(c)

	blog := models.NewBlog()
	category := models.NewCategory()
	categories := category.All()
	blogs := blog.FindListOrderCategory()

	this.SetResponse("BlogList", blogs)
	this.SetResponse("Categories", categories)

	this.MetaTitle = "TKO技術ブログ|カテゴリ別ブログ一覧"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "カテゴリ別ブログ一覧"
	this.MetaRobots = "noydir,noodp,index,follow"

	return this.Render(c, http.StatusOK, "category.html")
}
