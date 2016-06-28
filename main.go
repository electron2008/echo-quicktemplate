package main

import (
	"github.com/abenz1267/echo-quicktemplate/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Gzip())

	e.GET("/:name", controllers.Hello)
	e.GET("/", controllers.Index)
	e.Run(standard.New(":8080"))
}
