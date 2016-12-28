package controllers

import (
	"fmt"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"net/http"
)

type TopResponse struct {
	Blogs []*models.Blog
}

type TopController struct {
}

func (this TopController) Index(c echo.Context) error {
	blog := models.NewBlog()
	blogs := blog.FindList(10)

	data := new(TopResponse)
	data.Blogs = blogs

	req := c.Request()
	fmt.Printf("\n%v\n", req)
	fmt.Printf("%s", "==============")
	fmt.Printf("\n%v\n", req.Header)

	if c.QueryParam("ajax") == "on" {
		return c.JSON(http.StatusOK, map[string]interface{}{"Data": data})
	} else {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Data": data})
	}
}
