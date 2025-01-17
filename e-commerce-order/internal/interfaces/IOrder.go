package interfaces

import (
	"context"
	"e-commerce-order/external"
	"e-commerce-order/internal/models"
	"github.com/labstack/echo/v4"
)

type IOrderRepository interface {
	InsertNewOrder(ctx context.Context, order *models.Order) error
	UpdateStatusOrder(ctx context.Context, orderID int, status string) error
	GetDetailOrder(ctx context.Context, orderID int) (*models.Order, error)
	GetOrder(ctx context.Context) ([]models.Order, error)
}

type IOrderService interface {
	CreateOrder(ctx context.Context, profile *external.Profile, req *models.Order) (*models.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID int, req *models.OrderStatusRequest) error
	GetOrderList(ctx context.Context) ([]models.Order, error)
	GetOrderDetail(ctx context.Context, orderID int) (*models.Order, error)
}

type IOrderAPI interface {
	CreateOrder(e echo.Context) error
	UpdateOrderStatus(e echo.Context) error
	GetOrderDetail(e echo.Context) error
	GetOrderList(e echo.Context) error
}
