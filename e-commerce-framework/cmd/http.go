package cmd

import (
	"e-commerce-framework/helpers"
	"e-commerce-framework/internal/api"
	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	var (
		log = helpers.Logger
	)

	healthCheckAPI := &api.HealthCheckAPI{}

	e := echo.New()
	e.GET("/healthcheck", healthCheckAPI.HealthCheck)

	err := e.Start(":" + helpers.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}
