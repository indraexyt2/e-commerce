package interfaces

import (
	"context"
	"e-commerce-payment/internal/models"
	"github.com/labstack/echo/v4"
)

type IPaymentRepository interface {
	InsertNewPaymentMethod(ctx context.Context, req *models.PaymentMethod) error
	DeletePaymentMethod(ctx context.Context, sourceID int, userID int, sourceName string) error
	GetPaymentMethod(ctx context.Context, userID int, sourceName string) (*models.PaymentMethod, error)
	InsertNewPaymentTransaction(ctx context.Context, req *models.PaymentTransaction) error
	GetPaymentByOrderID(ctx context.Context, orderID int) (*models.PaymentTransaction, error)
	GetPaymentMethodById(ctx context.Context, paymentMethodID int) (*models.PaymentMethod, error)
	InsertNewPaymentRefund(ctx context.Context, req *models.PaymentRefund) error
}

type IPaymentService interface {
	PaymentMethodLink(ctx context.Context, req *models.PaymentMethodLink) error
	PaymentMethodLinkConfirm(ctx context.Context, userID int, req *models.PaymentMethodLinkConfirm) error
	PaymentMethodUnlink(ctx context.Context, userID int, req *models.PaymentMethodLink) error
	InitiatePayment(ctx context.Context, req *models.PaymentInitiatePayload) error
	RefundPayment(ctx context.Context, req *models.RefundPayload) error
}

type IPaymentAPI interface {
	PaymentMethodLink(e echo.Context) error
	PaymentMethodLinkConfirm(e echo.Context) error
	PaymentMethodUnlink(e echo.Context) error

	InitiatePayment(kafkaPayload []byte) error
	RefundPayment(kafkaPayload []byte) error
}
