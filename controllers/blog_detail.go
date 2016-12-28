package controllers

import (
	"fmt"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type BlogDetailController struct {
}

func (this BlogDetailController) Detail(c echo.Context) error {

	blog_id, err := strconv.Atoi(c.Param("blog_id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	blog := models.NewBlog()
	blog.FindById(blog_id)

	return c.Render(http.StatusOK, "blog_detail.html", map[string]interface{}{"Blog": blog})
}
