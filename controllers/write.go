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

	this.MetaTitle = "tko blogs|ブログ詳細"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "ブログ登録"
	this.MetaRobots = "noydir,noodp,noindex,nofollow"

	category := models.NewCategory()
	categories := category.All()

	this.SetResponse("Categories", categories)

	return this.Render(c, http.StatusOK, "write.html")
}

func (this *WriteController) Edit(c echo.Context) error {

	this.MetaTitle = "TKO技術ブログ|ブログ詳細"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "ブログ編集"
	this.MetaRobots = "noydir,noodp,noindex,nofollow"

	category := models.NewCategory()
	categories := category.All()

	this.SetResponse("Categories", categories)

	blogId, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	blog.FindById(blogId)

	this.SetResponse("Blog", blog)

	return this.Render(c, http.StatusOK, "write.html")
}

func (this *WriteController) Regist(c echo.Context) error {

	releaseDate, err := time.Parse("2006/01/02", c.FormValue("release_date"))

	if err != nil {
		panic(err)
	}

	blogId := c.FormValue("blog_id")
	body := c.FormValue("body")
	keywords := c.FormValue("keywords")
	category, err := strconv.Atoi(c.FormValue("category"))

	if err != nil {
		return err
	}

	blog := models.Blog{Title: c.FormValue("title"), Body: string(body), Keywords: string(keywords), Category: uint(category), IsShow: true, ReleaseDate: releaseDate}

	if blogId != "" {
		intid, err := strconv.Atoi(blogId)

		if err != nil {
			return err
		}

		blog.ID = uint(intid)

		blog.Update()
	} else {

		blog.Regist()
	}

	blog_count := models.BlogCount{BlogID: blog.ID}
	blog_count.Regist()

	return c.JSON(http.StatusOK, "success")
}
