package main

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
)

func RoutesRegister(e *echo.Echo) {
	e.GET("/", Top)
	e.GET("/write_blog", WriteBlog)
	e.POST("/send_blog", SendBlog)
}

type ResponseBase struct {
	PageKey string
}

type TopResponse struct {
	ResponseBase
	Blogs []*models.Blog
}

func Top(c echo.Context) error {
	blog := models.NewBlog()
	blogs := blog.FindAll()

	data := new(TopResponse)
	data.Blogs = blogs
	data.PageKey = "TOP"

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Data": data})
}

func WriteBlog(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}

func SendBlog(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}
