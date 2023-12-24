package main

import (
	"gotutor/handlers"
	"gotutor/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Renderer = utils.NewTemplateRenderer()

	e.GET("/", handlers.IndexHandler)
	e.GET("/test", handlers.TestHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
