package handlers

import (
	"gotutor/utils"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	foo := utils.GetConfig().Foo

	return c.String(200, "home page "+foo)
}
