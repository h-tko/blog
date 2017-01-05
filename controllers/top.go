package controllers

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
)

type TopController struct {
	BaseController
}

func (this *TopController) Index(c echo.Context) error {
	blog := models.NewBlog()
	blogs := blog.FindList(10)

	this.SetResponse("BlogList", blogs)

	this.MetaTitle = "TKO技術ブログ|ブログ一覧"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "ブログ一覧"
	this.MetaRobots = "noydir,noodp,index,follow"

	return this.Render(c, http.StatusOK, "blog_list.html")
}
