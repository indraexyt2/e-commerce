package repository

import (
	"context"
	"e-commerce-order/internal/models"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) InsertNewOrder(ctx context.Context, order *models.Order) error {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(order).Error
		if err != nil {
			return err
		}

		for i, orderItem := range order.OrderItems {
			orderItem.OrderID = order.ID
			err := tx.Create(&orderItem).Error
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to create order item %v", orderItem))
			}
			order.OrderItems[i].ID = orderItem.ID
			order.OrderItems[i].OrderID = order.ID
		}
		return nil
	})

	return err
}

func (r *OrderRepository) UpdateStatusOrder(ctx context.Context, orderID int, status string) error {
	return r.DB.WithContext(ctx).Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
