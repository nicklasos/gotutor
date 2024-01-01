package routes

import (
	"gotutor/di"

	"github.com/labstack/echo/v4"
)

func SetupWebHandlers(e *echo.Echo) {

	home := di.InitHomeHandler("Home")
	test := di.InitTestHandler()

	e.GET("/", home.Index)
	e.GET("/test", test.Index)
}
