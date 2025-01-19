package repository

import (
	"context"
	"e-commerce-payment/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func (r *PaymentRepository) InsertNewPaymentMethod(ctx context.Context, req *models.PaymentMethod) error {
	return r.DB.WithContext(ctx).Create(req).Error
}

func (r *PaymentRepository) DeletePaymentMethod(ctx context.Context, sourceID int, userID int, sourceName string) error {
	return r.DB.WithContext(ctx).Where("user_id = ?", userID).Where("source_name = ?", "e-commerce").Where("source_id = ?", sourceID).Where("source_name = ?", sourceName).Delete(&models.PaymentMethod{}).Error
}
