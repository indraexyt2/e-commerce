package cmd

import (
	"e-commerce-payment/external"
	"e-commerce-payment/helpers"
	"e-commerce-payment/internal/api"
	"e-commerce-payment/internal/interfaces"
	"e-commerce-payment/internal/repository"
	"e-commerce-payment/internal/services"
	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	var (
		log = helpers.Logger
	)

	d := DependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealthCheckAPI.HealthCheck)

	paymentV1 := e.Group("/payment/v1")
	paymentV1.POST("/link", d.PaymentAPI.PaymentMethodLink, d.MiddlewareValidateAuth)
	paymentV1.POST("/link/confirm", d.PaymentAPI.PaymentMethodLinkConfirm, d.MiddlewareValidateAuth)
	paymentV1.POST("/unlink", d.PaymentAPI.PaymentMethodUnlink, d.MiddlewareValidateAuth)

	err := e.Start(":" + helpers.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}

type Dependency struct {
	HealthCheckAPI *api.HealthCheckAPI
	External       interfaces.IExternal
	PaymentAPI     interfaces.IPaymentAPI
}

func DependencyInject() *Dependency {

	paymentRepo := &repository.PaymentRepository{DB: helpers.DB}
	paymentSvc := &services.PaymentService{
		PaymentRepository: paymentRepo,
		External:          &external.External{},
	}
	paymentApi := &api.PaymentAPI{PaymentService: paymentSvc}

	return &Dependency{
		HealthCheckAPI: &api.HealthCheckAPI{},
		External:       &external.External{},

		PaymentAPI: paymentApi,
	}
}
