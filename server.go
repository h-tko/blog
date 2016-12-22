package main

import (
	"fmt"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
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

	models.InitDatabase()
	defer models.CloseDatabase()

	e := echo.New()
	e.Static("/static", "assets")

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	RoutesRegister(e)

	config, err := toml.LoadFile("./config/app.toml")

	if err != nil {
		panic(err)
	}

	port := config.Get("application.port").(string)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
