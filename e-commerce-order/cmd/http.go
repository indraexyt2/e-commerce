package cmd

import (
	"e-commerce-order/external"
	"e-commerce-order/helpers"
	"e-commerce-order/internal/api"
	"e-commerce-order/internal/interfaces"
	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	var (
		log = helpers.Logger
	)

	d := DependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealthCheckAPI.HealthCheck)

	err := e.Start(":" + helpers.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}

type Dependency struct {
	HealthCheckAPI *api.HealthCheckAPI
	External       interfaces.IExternal
}

func DependencyInject() *Dependency {
	return &Dependency{
		HealthCheckAPI: &api.HealthCheckAPI{},
		External:       &external.External{},
	}
}
