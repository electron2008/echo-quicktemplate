package controllers

import (
	"github.com/abenz1267/echo-quicktemplate/templates"
	"github.com/labstack/echo"
)

// Renders template as response to Echo
func Render(c echo.Context, p templates.Page) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	templates.WritePageTemplate(c.Response(), p)
	return nil
}
