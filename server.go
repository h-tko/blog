package main

import (
	"fmt"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pelletier/go-toml"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	println("start application.")

	models.InitDatabase()
	defer models.CloseDatabase()

	println("Echo framework build.")

	e := echo.New()
	e.Use(middleware.CSRF())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	println("regist static dir.")
	e.Static("/static", "assets")

	t := &Template{
		templates: template.Must(template.New("").Funcs(TemplateHelpers).ParseGlob("views/*.html")),
	}

	e.Renderer = t

	println("regist routing.")
	routes(e)

	config, err := toml.LoadFile("./config/app.toml")

	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	port := config.Get("application.port").(string)

	println("start server.")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
