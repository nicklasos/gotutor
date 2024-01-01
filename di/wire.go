//go:build wireinject
// +build wireinject

package di

import (
	"gotutor/db"
	"gotutor/handlers"
	"gotutor/services"

	"github.com/google/wire"
)

func InitializeHomeHandler(name string) *handlers.HomeHandler {
	wire.Build(handlers.NewHomeHandler, services.NewMenuService)
	return &handlers.HomeHandler{}
}

func InitializeTestHandler() *handlers.TestHandler {
	wire.Build(handlers.NewTestHandler, services.NewMenuService, db.NewDB)
	return &handlers.TestHandler{}
}
