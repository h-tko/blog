package main

import (
	"html/template"
)

var TemplateHelpers = template.FuncMap{
	"raw":       htmlRaw,
	"substring": substring,
}

func htmlRaw(html string) template.HTML {
	return template.HTML(html)
}

func substring(text string, from int, to int) string {
	return string(text[from:to])
}
