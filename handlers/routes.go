package handlers

import (
	"github.com/labstack/echo/v4"
)

func SetupWebHandlers(e *echo.Echo) {

	home := NewHomeHandler()
	test := NewTestHandler()

	e.GET("/", home.Index)
	e.GET("/test", test.Index)
}
