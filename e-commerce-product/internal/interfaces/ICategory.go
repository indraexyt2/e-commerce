package interfaces

import (
	"context"
	"e-commerce-product/internal/models"
	"github.com/labstack/echo/v4"
)

type ICategoryRepository interface {
	InsertNewCategory(ctx context.Context, category *models.ProductCategory) error
}

type ICategoryService interface {
	CreateCategory(ctx context.Context, req *models.ProductCategory) (*models.ProductCategory, error)
}

type ICategoryAPI interface {
	CreateCategory(e echo.Context) error
}
