package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Index(c echo.Context) error {
	data := map[string]interface{}{
		"name": "Nick!",
	}

	return c.Render(http.StatusOK, "index.html", data)
}
