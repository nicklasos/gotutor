package handlers

import (
	"gotutor/db"
	"gotutor/services"
	"gotutor/views"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TestHandler struct {
	db   *gorm.DB
	menu *services.MenuService
}

func NewTestHandler(db *gorm.DB, menu *services.MenuService) *TestHandler {
	return &TestHandler{db, menu}
}

func (h *TestHandler) Index(c echo.Context) error {
	var user db.User
	if err := h.db.First(&user, 3).Error; err != nil {
		return views.Render(c, 500, views.ErrorHtml(err.Error()))
	}

	return views.Render(c, 200, views.Test())
}
