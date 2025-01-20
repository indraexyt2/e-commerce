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

func (r *PaymentRepository) GetPaymentMethod(ctx context.Context, userID int, sourceName string) (*models.PaymentMethod, error) {
	var (
		resp = &models.PaymentMethod{}
		err  error
	)
	err = r.DB.Where("user_id = ?", userID).Where("source_name = ?", sourceName).First(resp).Error
	return resp, err
}

func (r *PaymentRepository) GetPaymentMethodById(ctx context.Context, paymentMethodID int) (*models.PaymentMethod, error) {
	var (
		resp = &models.PaymentMethod{}
		err  error
	)
	err = r.DB.WithContext(ctx).Where("id = ?", paymentMethodID).First(resp).Error
	return resp, err
}

func (r *PaymentRepository) InsertNewPaymentTransaction(ctx context.Context, req *models.PaymentTransaction) error {
	return r.DB.WithContext(ctx).Create(req).Error
}

func (r *PaymentRepository) InsertNewPaymentRefund(ctx context.Context, req *models.PaymentRefund) error {
	return r.DB.WithContext(ctx).Create(req).Error
}

func (r *PaymentRepository) GetPaymentByOrderID(ctx context.Context, orderID int) (*models.PaymentTransaction, error) {
	var (
		resp = &models.PaymentTransaction{}
		err  error
	)
	err = r.DB.WithContext(ctx).Where("order_id = ?", orderID).First(resp).Error
	return resp, err
}
