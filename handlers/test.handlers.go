package handlers

import (
	"fmt"
	"gotutor/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TestHandler struct {
	Name string
}

func NewTestHandler(name string) *TestHandler {
	return &TestHandler{Name: name}
}

func (h *TestHandler) Index(c echo.Context) error {
	data := views.Data{
		"name": h.Name,
	}

	fmt.Println(h.Name)

	return c.Render(http.StatusOK, "index.html", data)
}
