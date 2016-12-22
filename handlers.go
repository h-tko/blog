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

func Top(c echo.Context) error {
	blog := models.NewBlog()
	blogs := blog.FindAll()
	print("%v", blogs)

	return c.Render(http.StatusOK, "index.html", "")
}

func WriteBlog(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}

func SendBlog(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}
