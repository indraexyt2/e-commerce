package cmd

import (
	"e-commerce-ums/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)
		auth := e.Request().Header.Get("Authorization")
		if auth == "" {
			log.Println("authorization header is empty")
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		_, err := d.UserRepository.GetUserSessionByToken(e.Request().Context(), auth)
		if err != nil {
			log.Println("failed to get user session by token: ", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		claim, err := helpers.ValidateToken(e.Request().Context(), auth)
		if err != nil {
			log.Println("failed to validate token: ", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println("token is expired: ", claim.ExpiresAt)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		e.Set("token", claim)
		return next(e)
	}
}

func (d *Dependency) MiddlewareRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)
		auth := e.Request().Header.Get("Authorization")
		if auth == "" {
			log.Println("authorization header is empty")
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		_, err := d.UserRepository.GetUserSessionByRefreshToken(e.Request().Context(), auth)
		if err != nil {
			log.Println("failed to get user session by token: ", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		claim, err := helpers.ValidateToken(e.Request().Context(), auth)
		if err != nil {
			log.Println("failed to validate token: ", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println("token is expired: ", claim.ExpiresAt)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		e.Set("token", claim)
		return next(e)
	}
}
