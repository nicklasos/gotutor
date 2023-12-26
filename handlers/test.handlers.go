package handlers

import (
	"gotutor/views"

	"github.com/labstack/echo/v4"
)

type TestHandler struct {
	Name string
}

func NewTestHandler(name string) *TestHandler {
	return &TestHandler{Name: name}
}

func (h *TestHandler) Index(c echo.Context) error {
	return views.Render(c, 200, views.Test(h.Name))
}
