package controllers

import (
	"net/http"

	"github.com/abenz1267/echo-quicktemplate/templates"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	// template to write
	p := &templates.Main{}

	// renders the template as a response
	return c.HTML(http.StatusOK, templates.PageTemplate(p))
}
