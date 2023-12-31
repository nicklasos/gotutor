//go:build wireinject
// +build wireinject

package di

import (
	"gotutor/handlers"
	"gotutor/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeHomeHandler(name string) *handlers.HomeHandler {
	wire.Build(handlers.NewHomeHandler, services.NewMenuService)
	return &handlers.HomeHandler{}
}

func InitializeTestHandler(db *gorm.DB) *handlers.TestHandler {
	wire.Build(handlers.NewTestHandler)
	return &handlers.TestHandler{}
}
