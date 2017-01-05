package controllers

import (
	"github.com/h-tko/blog/libraries"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type WriteController struct {
	BaseController
}

func (this *WriteController) Write(c echo.Context) error {

	this.MetaTitle = "TKO技術ブログ|ブログ詳細"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "ブログ登録"
	this.MetaRobots = "noydir,noodp,noindex,nofollow"

	return this.Render(c, http.StatusOK, "write.html")
}

func (this *WriteController) Regist(c echo.Context) error {

	release_date, err := time.Parse("2006/01/02", c.FormValue("release_date"))

	if err != nil {
		panic(err)
	}

	body := libraries.MarkdownToHtml(c.FormValue("body"))

	blog := models.Blog{Title: c.FormValue("title"), Body: string(body), IsShow: true, ReleaseDate: release_date}
	models.RegistBlog(&blog)

	blog_count := models.BlogCount{BlogID: blog.ID}
	models.RegistBlogCount(&blog_count)

	return c.JSON(http.StatusOK, "success")
}
