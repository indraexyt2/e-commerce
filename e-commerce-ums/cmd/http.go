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
	userV1.POST("/register/admin", d.UserAPI.RegisterAdmin)
	userV1.POST("/login", d.UserAPI.LoginUser)
	userV1.POST("/login/admin", d.UserAPI.LoginAdmin)
	userV1.GET("/profile", d.UserAPI.GetProfile, d.MiddlewareValidateAuth)
	userV1.PUT("/refresh-token", d.UserAPI.RefreshToken, d.MiddlewareRefreshToken)

	err := e.Start(":" + helpers.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthCheckAPI *api.HealthCheckAPI
	UserAPI        interfaces.IUserAPI
}

func DependencyInject() *Dependency {
	userRepo := &repository.UserRepository{DB: helpers.DB}

	userSvc := &services.UserService{UserRepository: userRepo}
	userApi := &api.UserAPI{UserService: userSvc}

	return &Dependency{
		UserRepository: userRepo,
		HealthCheckAPI: &api.HealthCheckAPI{},
		UserAPI:        userApi,
	}
}
