package cmd

import (
	"e-commerce-order/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)

		auth := e.Request().Header.Get("Authorization")
		if auth == "" {
			log.Error("authorization header is empty")
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		profile, err := d.External.GetProfile(e.Request().Context(), auth)
		if err != nil || profile == nil {
			log.Error("failed to get profile: ", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		e.Set("profile", profile)
		return next(e)
	}
}
