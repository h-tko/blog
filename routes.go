package main

import (
	. "github.com/h-tko/blog/controllers"
	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/", new(TopController).Index)
	e.GET("/detail/:blog_id/", new(BlogDetailController).Detail)
	e.GET("/write_blog/", new(WriteController).Write)
	e.POST("/regist_blog/", new(WriteController).Regist)
	e.GET("/increment_good/", new(BlogDetailController).IncrementGood)
	e.GET("/popular_list/", new(RightColumnController).PopularList)
	e.GET("/edit_blog/:blog_id/", new(WriteController).Edit)
	e.GET("/category/", new(CategoryController).Index)
	e.GET("/category/:category_id/", new(CategoryController).One)
}
