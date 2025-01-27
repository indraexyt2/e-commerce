package interfaces

import (
	"e-commerce-order/external"
	"golang.org/x/net/context"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (*external.Profile, error)
	ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error
}
