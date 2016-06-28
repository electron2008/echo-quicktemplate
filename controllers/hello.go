package controllers

import (
	"net/http"

	"github.com/abenz1267/echo-quicktemplate/templates"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	// initialize a map named "vars". This map will hold all the variables for the template
	vars := make(map[string]interface{})
	vars["title"] = "Hello-World"
	vars["name"] = c.Param("name")

	// Template to render + map "vars" to struct-field "Vars"
	p := &templates.Hello{
		Vars: vars,
	}

	return c.HTML(http.StatusOK, templates.PageTemplate(p))
}
