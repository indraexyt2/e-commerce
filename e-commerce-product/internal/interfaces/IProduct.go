package interfaces

import (
	"context"
	"e-commerce-product/internal/models"
	"github.com/labstack/echo/v4"
)

type IProductRepository interface {
	InsertNewProduct(ctx context.Context, product *models.Product) error
}

type IProductService interface {
	CreateProduct(ctx context.Context, req *models.Product) (*models.Product, error)
}

type IProductAPI interface {
	CreateProduct(e echo.Context) error
}
