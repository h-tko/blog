package main

import (
	"fmt"
	"github.com/echo-contrib/sessions"
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pelletier/go-toml"
	"html/template"
	"io"
	"strings"
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
	//    e.Use(middleware.CSRF())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config, err := toml.LoadFile("./config/app.toml")

	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	cache_host := config.Get("cache.host").(string)
	cache_port := config.Get("cache.port").(string)
	cache_secret := config.Get("cache.secret").(string)

	store, err := sessions.NewRedisStore(16, "tcp", fmt.Sprintf("%s:%s", cache_host, cache_port), "", []byte(cache_secret))

	if err != nil {
		panic(err)
	}

	e.Use(sessions.Sessions("echosession", store))

	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Index(c.Request().RequestURI, "/write_blog/") >= 0 {
				return false
			}

			return true
		},
		Validator: func(username, password string, c echo.Context) bool {
			if username == "h-tko" && password == "WeJLn0mW" {
				return true
			}

			return false
		},
	}))
	e.Pre(middleware.AddTrailingSlash())

	println("regist static dir.")
	e.Static("/static", "assets")

	t := &Template{
		templates: template.Must(template.New("").Funcs(TemplateHelpers).ParseGlob("views/*.html")),
	}

	e.Renderer = t

	println("regist routing.")
	routes(e)

	port := config.Get("application.port").(string)

	println("start server.")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
