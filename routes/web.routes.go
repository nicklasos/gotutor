package routes

import (
	"gotutor/di"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupWebHandlers(e *echo.Echo, db *gorm.DB) {

	home := di.InitializeHomeHandler("Home")
	test := di.InitializeTestHandler(db)

	e.GET("/", home.Index)
	e.GET("/test", test.Index)
}
