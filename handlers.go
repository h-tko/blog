package main

import (
	"github.com/h-tko/blog/libraries"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func RoutesRegister(e *echo.Echo) {
	e.GET("/", top)
	e.GET("/write_blog", writeBlog)
	e.POST("/regist_blog", registBlog)
}

type ResponseBase struct {
	PageKey string
}

type TopResponse struct {
	ResponseBase
	Blogs []*models.Blog
}

func top(c echo.Context) error {
	blog := models.NewBlog()
	blogs := blog.FindList(10)

	data := new(TopResponse)
	data.Blogs = blogs
	data.PageKey = "TOP"

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Data": data})
}

func writeBlog(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Data": nil})
}

func registBlog(c echo.Context) error {

	release_date, err := time.Parse("2006/01/02", c.FormValue("release_date"))

	if err != nil {
		panic(err)
	}

	body := libraries.MarkdownToHtml(c.FormValue("body"))

	blog := models.Blog{Title: c.FormValue("title"), Body: string(body), IsShow: c.FormValue("is_show") == "1", ReleaseDate: release_date}
	models.RegistBlog(blog)

	return c.JSON(http.StatusOK, blog)
}
