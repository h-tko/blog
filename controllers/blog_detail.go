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

	blog_id, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	blog.FindById(blog_id)

	this.SetResponse("Blog", blog)

	this.MetaTitle = "TKO技術ブログ|ブログ詳細"
	this.MetaDescription = "TKO技術ブログです"
	this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	this.MetaH1 = "ブログ詳細"
	this.MetaRobots = "noydir,noodp,index,follow"

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
