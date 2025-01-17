package services

import (
	"context"
	"e-commerce-order/constants"
	"e-commerce-order/external"
	"e-commerce-order/helpers"
	"e-commerce-order/internal/interfaces"
	"e-commerce-order/internal/models"
	"encoding/json"
	"github.com/pkg/errors"
)

type OrderService struct {
	OrderRepository interfaces.IOrderRepository
	External        interfaces.IExternal
}

func (s *OrderService) CreateOrder(ctx context.Context, profile *external.Profile, req *models.Order) (*models.Order, error) {
	req.UserID = profile.Data.UserID
	req.Status = constants.OrderStatusPending
	err := s.OrderRepository.InsertNewOrder(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create order")
	}

	// produce new message
	kafkaPayload := models.PaymentInitiatePayload{
		UserID:     profile.Data.UserID,
		OrderID:    req.ID,
		TotalPrice: req.TotalPrice,
	}
	jsonPayload, err := json.Marshal(kafkaPayload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal kafka payload")
	}
	kafkaErr := s.External.ProduceKafkaMessage(ctx, helpers.GetEnv("KAFKA_TOPIC_PAYMENT_INITIATE"), jsonPayload)
	if kafkaErr != nil {
		errUpdateStatus := s.OrderRepository.UpdateStatusOrder(ctx, req.ID, constants.OrderStatusFailed)
		if errUpdateStatus != nil {
			helpers.Logger.Error("Error updating order status to failed: ", errUpdateStatus)
		}
		return nil, errors.Wrap(kafkaErr, "failed to produce kafka message")
	}

	return req, nil
}
