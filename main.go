package main

import (
	"gotutor/routes"
	"gotutor/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Renderer = utils.NewTemplateRenderer()

	routes.InitWebHandlers(e)

	e.Logger.Fatal(e.Start(":8080"))
}