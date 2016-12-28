package controllers

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
)

type TopController struct {
}

func (this TopController) Index(c echo.Context) error {
	blog := models.NewBlog()
	blogs := blog.FindList(10)

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{"BlogList": blogs})
}
