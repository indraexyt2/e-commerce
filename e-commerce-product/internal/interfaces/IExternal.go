package interfaces

import (
	"context"
	"e-commerce-product/external"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (*external.Profile, error)
}
