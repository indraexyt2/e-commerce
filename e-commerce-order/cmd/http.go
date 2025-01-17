package cmd

import (
	"e-commerce-order/external"
	"e-commerce-order/helpers"
	"e-commerce-order/internal/api"
	"e-commerce-order/internal/interfaces"
	"e-commerce-order/internal/repository"
	"e-commerce-order/internal/services"
	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	var (
		log = helpers.Logger
	)

	d := DependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealthCheckAPI.HealthCheck)

	orderV1 := e.Group("/order/v1")
	orderV1.POST("", d.OrderAPI.CreateOrder, d.MiddlewareValidateAuth)
	orderV1.GET("/:id", d.OrderAPI.GetOrderDetail, d.MiddlewareValidateAuth)
	orderV1.GET("", d.OrderAPI.GetOrderList, d.MiddlewareValidateAuth)

	err := e.Start(":" + helpers.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}

type Dependency struct {
	HealthCheckAPI *api.HealthCheckAPI
	External       interfaces.IExternal

	OrderAPI interfaces.IOrderAPI
}

func DependencyInject() *Dependency {
	orderRepo := &repository.OrderRepository{DB: helpers.DB}
	orderSvc := &services.OrderService{
		OrderRepository: orderRepo,
		External:        &external.External{},
	}
	orderApi := &api.OrderAPI{OrderService: orderSvc}

	return &Dependency{
		HealthCheckAPI: &api.HealthCheckAPI{},
		External:       &external.External{},
		OrderAPI:       orderApi,
	}
}
