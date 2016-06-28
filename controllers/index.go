package controllers

import (
	"net/http"

	"github.com/abenz1267/echo-quicktemplate/templates"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	// template to write
	p := &templates.Main{}

	// converts buf to a string, so c.HTML can render it to the browser
	return c.HTML(http.StatusOK, templates.PageTemplate(p))
}
