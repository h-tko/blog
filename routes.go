package main

import (
	. "github.com/h-tko/blog/controllers"
	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/", TopController{}.Index)
	e.GET("/detail/:blog_id/", BlogDetailController{}.Detail)
	e.GET("/write_blog/", WriteController{}.Write)
	e.POST("/regist_blog/", WriteController{}.Regist)
}
