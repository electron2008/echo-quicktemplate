package controllers

import (
	"bytes"
	"net/http"

	"github.com/abenz1267/echo-quicktemplate/templates"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	// Buffer the template is going to be written to
	buf := new(bytes.Buffer)

	// template to write
	p := &templates.Main{}

	// writes the template to the buffer
	templates.WritePageTemplate(buf, p)

	// converts buf to a string, so c.HTML can render it to the browser
	return c.HTML(http.StatusOK, buf.String())
}
