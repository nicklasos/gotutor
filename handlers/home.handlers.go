package handlers

import (
	"gotutor/services"
	"gotutor/utils"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	menu services.MenuService
}

func NewHomeHandler(menu *services.MenuService) *HomeHandler {
	return &HomeHandler{
		menu: *menu,
	}
}

func (h *HomeHandler) Index(c echo.Context) error {
	foo := utils.GetConfig().Foo
	bar := h.menu.Name

	return c.String(200, "home page "+foo+" "+bar)
}
