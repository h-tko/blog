package main

import (
	"html/template"
)

var TemplateHelpers = template.FuncMap{
	"raw":      htmlRaw,
	"mbsubstr": mbsubstr,
}

func htmlRaw(html string) template.HTML {
	return template.HTML(html)
}

func mbsubstr(text string, from int, to int) string {

	rntext := []rune(text)

	if len(rntext) <= to {
		return text
	}

	return string(rntext[from:to])
}
