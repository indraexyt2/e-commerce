package interfaces

import (
	"e-commerce-payment/external"
	"e-commerce-payment/internal/models"
	"golang.org/x/net/context"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (*external.Profile, error)
	ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error

	PaymentLink(ctx context.Context, req *models.PaymentMethodLink) (*external.PaymentLinkResponse, error)
	PaymentUnlink(ctx context.Context, req *models.PaymentMethodLink) (*external.PaymentLinkResponse, error)
	PaymentLinkConfirmation(ctx context.Context, sourceID int, otp string) (*external.PaymentLinkResponse, error)
	WalletTransaction(ctx context.Context, req *external.PaymentTransactionRequest) (*external.PaymentTransactionResponse, error)

	OrderCallback(ctx context.Context, orderID int, status string) (*external.OrderResponse, error)
}
