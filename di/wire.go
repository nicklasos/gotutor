//go:build wireinject
// +build wireinject

package di

import (
	"gotutor/db"
	"gotutor/handlers"
	"gotutor/services"

	"github.com/google/wire"
)

func InitHomeHandler(name string) *handlers.HomeHandler {
	wire.Build(handlers.NewHomeHandler, services.NewMenuService)
	return &handlers.HomeHandler{}
}

func InitTestHandler() *handlers.TestHandler {
	wire.Build(handlers.NewTestHandler, services.NewMenuService, db.NewDB)
	return &handlers.TestHandler{}
}
