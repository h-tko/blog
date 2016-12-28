package controllers

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type BlogDetailController struct {
}

func (this BlogDetailController) Detail(c echo.Context) error {

	if c.QueryParam("ajax") != "on" {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	}

	blog_id, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		return err
	}

	blog := models.NewBlog()
	blog.FindById(blog_id)

	return c.JSON(http.StatusOK, map[string]interface{}{"blog": blog})
}
