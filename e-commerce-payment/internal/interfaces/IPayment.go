package interfaces

import (
	"context"
	"e-commerce-payment/internal/models"
	"github.com/labstack/echo/v4"
)

type IPaymentRepository interface {
	InsertNewPaymentMethod(ctx context.Context, req *models.PaymentMethod) error
	DeletePaymentMethod(ctx context.Context, sourceID int, userID int, sourceName string) error
}

type IPaymentService interface {
	PaymentMethodLink(ctx context.Context, req *models.PaymentMethodLink) error
	PaymentMethodLinkConfirm(ctx context.Context, userID int, req *models.PaymentMethodLinkConfirm) error
	PaymentMethodUnlink(ctx context.Context, userID int, req *models.PaymentMethodLink) error
}

type IPaymentAPI interface {
	PaymentMethodLink(e echo.Context) error
	PaymentMethodLinkConfirm(e echo.Context) error
	PaymentMethodUnlink(e echo.Context) error
}
