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
	return c.Render(http.StatusOK, "write.html", nil)
}

func (this WriteController) Regist(c echo.Context) error {

	release_date, err := time.Parse("2006/01/02", c.FormValue("release_date"))

	if err != nil {
		panic(err)
	}

	body := libraries.MarkdownToHtml(c.FormValue("body"))

	blog := models.Blog{Title: c.FormValue("title"), Body: string(body), IsShow: c.FormValue("is_show") == "1", ReleaseDate: release_date}
	models.RegistBlog(blog)

	blog_count := models.BlogCount{BlogId: blog.ID}
	models.RegistBlogCount(blog_count)

	return c.JSON(http.StatusOK, "success")
}
