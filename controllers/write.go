package controllers

import (
	"github.com/h-tko/blog/libraries"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type WriteController struct {
}

func (this WriteController) Write(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Data": nil})
}

func (this WriteController) Regist(c echo.Context) error {

	release_date, err := time.Parse("2006/01/02", c.FormValue("release_date"))

	if err != nil {
		panic(err)
	}

	body := libraries.MarkdownToHtml(c.FormValue("body"))

	blog := models.Blog{Title: c.FormValue("title"), Body: string(body), IsShow: c.FormValue("is_show") == "1", ReleaseDate: release_date}
	models.RegistBlog(blog)

	return c.JSON(http.StatusOK, "success")
}
