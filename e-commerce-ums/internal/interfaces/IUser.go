package interfaces

import (
	"context"
	"e-commerce-ums/internal/models"
	"github.com/labstack/echo/v4"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, req *models.User) (*models.User, error)
	RegisterAdmin(ctx context.Context, req *models.User) (*models.User, error)
}

type IUserAPI interface {
	RegisterUser(e echo.Context) error
	RegisterAdmin(e echo.Context) error
}
