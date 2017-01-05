package controllers

import (
	"github.com/labstack/echo"
)

type BaseController struct {
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	MetaH1          string
	MetaRobots      string

	response map[string]interface{}
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

	this.response["mt"] = this.MetaTitle
	this.response["md"] = this.MetaDescription
	this.response["mk"] = this.MetaKeywords
	this.response["mh1"] = this.MetaH1
	this.response["mr"] = this.MetaRobots

	return c.Render(status, oFile, this.response)
}
