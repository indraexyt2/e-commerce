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
}

type IuserAPI interface {
	RegisterUser(e echo.Context) error
}
