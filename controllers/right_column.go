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
	popularList := model.FindPopularList(5)

	return c.JSON(http.StatusOK, map[string]interface{}{"PopularList": popularList})
}

func (this *RightColumnController) CategoryChart(c echo.Context) error {
	model := models.NewBlog()
	categoryCounts := model.CategoryCount()

	return c.JSON(http.StatusOK, map[string]interface{}{"CategoryCounts": categoryCounts})
}
