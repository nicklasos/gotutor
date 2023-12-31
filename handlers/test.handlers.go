package handlers

import (
	"gotutor/db"
	"gotutor/views"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TestHandler struct {
	db *gorm.DB
}

func NewTestHandler(db *gorm.DB) *TestHandler {
	return &TestHandler{db: db}
}

func (h *TestHandler) Index(c echo.Context) error {
	var user db.User

	if err := h.db.First(&user, 3).Error; err != nil {
		return views.Render(c, 500, views.ErrorHtml(err.Error()))
	}

	return views.Render(c, 200, views.Test())
}
