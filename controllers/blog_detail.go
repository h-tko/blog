package controllers

import (
	"fmt"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type BlogDetailController struct {
	BaseController
}

func (this *BlogDetailController) Detail(c echo.Context) error {
	this.BeforeFilter(c)

	blog_id, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	blog.FindById(blog_id)

	this.SetResponse("Blog", blog)

	this.MetaTitle = fmt.Sprintf("TKO技術ブログ|%s", blog.Title)
	this.MetaDescription = "TKO技術ブログです"
	this.MetaH1 = fmt.Sprintf("ブログ詳細(%s)", blog.Title)
	this.MetaRobots = "noydir,noodp,index,follow"

	if len(blog.Keywords) > 0 {
		this.MetaKeywords = fmt.Sprintf("%s,テックブログ,技術ブログ,IT,ブログ", blog.Keywords)
	} else {
		this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	}

	return this.Render(c, http.StatusOK, "blog_detail.html")
}

func (this *BlogDetailController) IncrementGood(c echo.Context) error {

	blog_id, err := strconv.Atoi(c.QueryParam("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog_count := models.NewBlogCount()
	result_count := blog_count.IncrementGood(uint(blog_id))

	return c.JSON(http.StatusOK, result_count)
}
