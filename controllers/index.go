package controllers

import (
	"github.com/abenz1267/echo-quicktemplate/templates"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	// template to write
	p := &templates.Main{}

	// renders the template as a response
	// renders the template as a response
	return Render(c, p)
}
