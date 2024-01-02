package handlers

import (
	"gotutor/services"
	"gotutor/views"

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
	foo := "Foo"
	bar := "Bar"

	return views.Render(c, 200, views.Home(foo, bar))
}
