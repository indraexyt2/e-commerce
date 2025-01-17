package repository

import (
	"context"
	"e-commerce-order/internal/models"
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

		return nil
	})

	return err
}

func (r *OrderRepository) UpdateStatusOrder(ctx context.Context, orderID int, status string) error {
	return r.DB.WithContext(ctx).Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *OrderRepository) GetDetailOrder(ctx context.Context, orderID int) (*models.Order, error) {
	var (
		order = &models.Order{}
		err   error
	)
	err = r.DB.WithContext(ctx).Preload("OrderItems").Where("id = ?", orderID).First(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) GetOrder(ctx context.Context) ([]models.Order, error) {
	var (
		order []models.Order
		err   error
	)
	err = r.DB.WithContext(ctx).Model(&models.Order{}).Preload("OrderItems").Order("id DESC").Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
