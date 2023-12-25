package handlers

import (
	"gotutor/utils"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Index(c echo.Context) error {
	foo := utils.GetConfig().Foo

	return c.String(200, "home page "+foo)
}
