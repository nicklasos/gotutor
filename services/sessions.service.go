package services

import (
	"gotutor/utils"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func InitSessions(e *echo.Echo) {
	config := utils.GetConfig()
	key := config.APP_SECRET
	maxAge := 86400 * 30

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = config.APP_ENV == "production"
	gothic.Store = store
}
