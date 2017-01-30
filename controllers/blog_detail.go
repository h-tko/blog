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

	blogID, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	blog.FindByID(blogID)

	goodBlogIDs := session.Get("good_blog_ids")
	voted := false

	if goodBlogIDs != nil {
		for _, id := range goodBlogIDs.([]int) {
			if id == blogID {
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

	blogID, err := strconv.Atoi(c.QueryParam("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blogCount := models.NewBlogCount()
	resultCount := blogCount.IncrementGood(uint(blogID))

	ids := session.Get("good_blog_ids")
	var goodBlogIDs = []int{}

	if ids != nil {
		goodBlogIDs = append(ids.([]int), blogID)
	} else {
		goodBlogIDs = []int{blogID}
	}

	session.Set("good_blog_ids", goodBlogIDs)
	session.Save()

	return c.JSON(http.StatusOK, resultCount)
}
