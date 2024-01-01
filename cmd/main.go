package main

import (
	"gotutor/db"
	"gotutor/handlers"
	"gotutor/routes"
	"gotutor/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.InitConfig()

	db.InitDB()

	e := echo.New()

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	e.Use(middleware.Logger())

	routes.SetupWebHandlers(e)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(":8080"))
}
