package cmd

import (
	"e-commerce-product/external"
	"e-commerce-product/helpers"
	"e-commerce-product/internal/api"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/repository"
	"e-commerce-product/internal/services"
	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	var (
		log = helpers.Logger
	)

	d := DependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealthCheckAPI.HealthCheck)

	productV1 := e.Group("/product/v1")
	productV1.POST("", d.ProductAPI.CreateProduct, d.MiddlewareValidateAuth)

	categoryV1 := e.Group("/product/v1/category")
	categoryV1.POST("", d.CategoryAPI.CreateCategory, d.MiddlewareValidateAuth)

	err := e.Start(":" + helpers.GetEnv("PORT"))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}

type Dependency struct {
	HealthCheckAPI *api.HealthCheckAPI
	External       interfaces.IExternal

	ProductAPI  interfaces.IProductAPI
	CategoryAPI interfaces.ICategoryAPI
}

func DependencyInject() *Dependency {
	productRepo := &repository.ProductRepository{
		DB:    helpers.DB,
		Redis: helpers.RedisClient,
	}
	productSvc := &services.ProductService{ProductRepository: productRepo}
	productApi := &api.ProductAPI{ProductService: productSvc}

	categoryRepo := &repository.CategoryRepository{DB: helpers.DB}
	categorySvc := &services.CategoryService{CategoryRepository: categoryRepo}
	categoryApi := &api.CategoryAPI{CategoryService: categorySvc}

	return &Dependency{
		HealthCheckAPI: &api.HealthCheckAPI{},
		External:       &external.External{},

		ProductAPI:  productApi,
		CategoryAPI: categoryApi,
	}
}