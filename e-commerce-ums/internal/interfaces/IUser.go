package interfaces

import (
	"context"
	"e-commerce-ums/internal/models"
	"github.com/labstack/echo/v4"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string, role string) (*models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error)
}

type IUserService interface {
	Register(ctx context.Context, req *models.User, role string) (*models.User, error)
	Login(ctx context.Context, request *models.LoginRequest, role string) (*models.LoginResponse, error)
	GetProfile(ctx context.Context, username string) (*models.User, error)
}

type IUserAPI interface {
	RegisterUser(e echo.Context) error
	RegisterAdmin(e echo.Context) error
	LoginUser(e echo.Context) error
	LoginAdmin(e echo.Context) error
	GetProfile(e echo.Context) error
}
