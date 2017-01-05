package controllers

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
)

type RightColumnController struct {
	BaseController
}

func (this *RightColumnController) PopularList(c echo.Context) error {
	model := models.NewBlog()
	popular_list := model.FindPopularList(5)

	return c.JSON(http.StatusOK, map[string]interface{}{"PopularList": popular_list})
}
