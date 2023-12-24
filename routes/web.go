package routes

import (
	"gotutor/handlers"

	"github.com/labstack/echo/v4"
)

func InitWebHandlers(e *echo.Echo) {
	e.GET("/", handlers.IndexHandler)
	e.GET("/test", handlers.TestHandler)
}
