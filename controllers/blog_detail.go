package controllers

import (
	"fmt"
	"github.com/echo-contrib/sessions"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type BlogDetailController struct {
	BaseController
}

func (this *BlogDetailController) Detail(c echo.Context) error {
	session := sessions.Default(c)

	this.BeforeFilter(c)

	blog_id, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	blog.FindById(blog_id)

	good_blog_ids := session.Get("good_blog_ids")
	voted := false

	if good_blog_ids != nil {
		for _, id := range good_blog_ids.([]int) {
			if id == blog_id {
				voted = true
			}
		}
	}

	this.SetResponse("Voted", voted)
	this.SetResponse("Blog", blog)

	this.MetaTitle = fmt.Sprintf("tko blogs|%s", blog.Title)
	this.MetaDescription = "TKO技術ブログです"
	this.MetaH1 = fmt.Sprintf("%s", blog.Title)
	this.MetaRobots = "noydir,noodp,index,follow"

	if len(blog.Keywords) > 0 {
		this.MetaKeywords = fmt.Sprintf("%s,テックブログ,技術ブログ,IT,ブログ", blog.Keywords)
	} else {
		this.MetaKeywords = "テックブログ,技術ブログ,IT,ブログ"
	}

	return this.Render(c, http.StatusOK, "blog_detail.html")
}

func (this *BlogDetailController) IncrementGood(c echo.Context) error {
	session := sessions.Default(c)

	blog_id, err := strconv.Atoi(c.QueryParam("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog_count := models.NewBlogCount()
	result_count := blog_count.IncrementGood(uint(blog_id))

	ids := session.Get("good_blog_ids")
	var good_blog_ids = []int{}

	if ids != nil {
		good_blog_ids = append(ids.([]int), blog_id)
	} else {
		good_blog_ids = []int{blog_id}
	}

	session.Set("good_blog_ids", good_blog_ids)
	session.Save()

	return c.JSON(http.StatusOK, result_count)
}
