package controllers

import (
	"github.com/h-tko/blog/models"
	"github.com/labstack/echo"
	"strings"
)

type BaseController struct {
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	MetaH1          string
	MetaRobots      string

	response map[string]interface{}
}

func (this *BaseController) BeforeFilter(c echo.Context) {
	accessLog := models.NewAccessLog()
	request := c.Request()

	xForwardIp := request.Header.Get("X-Forwarded-For")

	if len(xForwardIp) > 0 {
		accessLog.IpAddress = xForwardIp
	} else {
		accessLog.IpAddress = request.RemoteAddr
	}

	accessLog.UserAgent = request.UserAgent()
	accessLog.Uri = request.RequestURI

	accessLog.Regist()
}

func (this *BaseController) SetResponse(key string, val interface{}) {

	if this.response == nil {
		this.response = make(map[string]interface{})
	}

	this.response[key] = val
}

func (this *BaseController) Render(c echo.Context, status int, oFile string) error {
	if this.response == nil {
		this.response = make(map[string]interface{})
	}

	if strings.Index(c.Request().RequestURI, "/category/") > -1 {
		this.SetResponse("tab", "category")
	} else {
		this.SetResponse("tab", "top")
	}

	this.response["mt"] = this.MetaTitle
	this.response["md"] = this.MetaDescription
	this.response["mk"] = this.MetaKeywords
	this.response["mh1"] = this.MetaH1
	this.response["mr"] = this.MetaRobots

	return c.Render(status, oFile, this.response)
}
