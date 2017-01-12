package controllers

import (
	"fmt"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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

func (this *WriteController) Edit(c echo.Context) error {

	this.MetaTitle = "TKO技術ブログ|ブログ詳細"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "ブログ編集"
	this.MetaRobots = "noydir,noodp,noindex,nofollow"

	blog_id, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	blog.FindById(blog_id)

	this.SetResponse("Blog", blog)

	return this.Render(c, http.StatusOK, "write.html")
}

func (this *WriteController) Regist(c echo.Context) error {

	release_date, err := time.Parse("2006/01/02", c.FormValue("release_date"))

	if err != nil {
		panic(err)
	}

	blog_id := c.FormValue("blog_id")
	body := c.FormValue("body")

	blog := models.Blog{Title: c.FormValue("title"), Body: string(body), IsShow: true, ReleaseDate: release_date}

	if blog_id != "" {
		intid, err := strconv.Atoi(blog_id)

		if err != nil {
			return err
		}

		blog.ID = uint(intid)

		models.UpdateBlog(&blog)
	} else {

		models.RegistBlog(&blog)
	}

	blog_count := models.BlogCount{BlogID: blog.ID}
	models.RegistBlogCount(&blog_count)

	return c.JSON(http.StatusOK, "success")
}
