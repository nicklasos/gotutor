package handlers

import (
	"context"
	"gotutor/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func InitAuthHandlers(e *echo.Echo) {
	config := utils.GetConfig()

	goth.UseProviders(
		google.New(
			config.GOOGLE_CLIENT_ID,
			config.GOOGLE_CLIENT_SECRET,
			config.GOOGLE_CALLBACK_URL,
			"email",
			"profile",
		),
	)

	e.GET("/auth/:provider", func(c echo.Context) error {
		provider := c.Param("provider")
		req := c.Request().WithContext(context.WithValue(context.Background(), "provider", provider))

		gothic.BeginAuthHandler(c.Response().Writer, req)

		return nil
	})

	e.GET("/auth/:provider/callback", func(c echo.Context) error {
		user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	})
}
