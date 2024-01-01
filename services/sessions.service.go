package services

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func InitSessions() {
	key := "session_secret"
	maxAge := 86400 * 30

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = false // Set to true in production
	gothic.Store = store
}
