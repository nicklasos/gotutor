package handlers

import (
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
	// var user db.User
	// if err := h.db.First(&user, 3).Error; err != nil {
	// 	return views.Render(c, 500, views.ErrorHtml(err.Error()))
	// }

	// sess, err := session.Get("session", c)
	// if err != nil {
	// 	return views.Render(c, 500, views.ErrorHtml(err.Error()))
	// }

	// sess.Values["foo"] = "bar"
	// sess.Save(c.Request(), c.Response())

	return views.Render(c, 200, views.Test())
}
