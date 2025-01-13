package cmd

import (
	"e-commerce-ums/helpers"
	"e-commerce-ums/internal/api"
	"e-commerce-ums/internal/interfaces"
	"e-commerce-ums/internal/repository"
	"e-commerce-ums/internal/services"
	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	var (
		log = helpers.Logger
	)

	d := DependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealthCheckAPI.HealthCheck)

	userV1 := e.Group("/user/v1")
	userV1.POST("/register", d.UserAPI.RegisterUser)

	err := e.Start(":" + helpers.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}

type Dependency struct {
	HealthCheckAPI *api.HealthCheckAPI

	UserAPI interfaces.IuserAPI
}

func DependencyInject() *Dependency {
	userRepo := &repository.UserRepository{DB: helpers.DB}

	userSvc := &services.UserService{UserRepository: userRepo}
	userApi := &api.UserAPI{UserService: userSvc}

	return &Dependency{
		HealthCheckAPI: &api.HealthCheckAPI{},
		UserAPI:        userApi,
	}
}
