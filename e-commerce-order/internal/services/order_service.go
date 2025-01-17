package services

import (
	"context"
	"e-commerce-order/constants"
	"e-commerce-order/external"
	"e-commerce-order/helpers"
	"e-commerce-order/internal/interfaces"
	"e-commerce-order/internal/models"
	"encoding/json"
	"fmt"
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

func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID int, profile *external.Profile, req *models.OrderStatusRequest) error {
	if !constants.MappingOrderStatus[req.Status] {
		return fmt.Errorf("invalid status: %v", req.Status)
	}

	order, err := s.OrderRepository.GetDetailOrder(ctx, orderID)
	if err != nil {
		return errors.Wrap(err, "failed to get detail order")
	}

	validStatusReq := false
	statusFlow := constants.MappingFlowOrderStatus[order.Status]
	for i := range statusFlow {
		if statusFlow[i] == req.Status {
			validStatusReq = true
			break
		}
	}

	if !validStatusReq {
		return fmt.Errorf("invalid status flow, current status: %v, request status: %v", order.Status, req.Status)
	}

	if req.Status == constants.OrderStatusRefund {
		if profile.Data.Role != "admin" {
			return errors.New("only admin can refund order")
		}

		kafkaPayload := &models.RefundPayload{
			OrderID: order.ID,
			AdminID: profile.Data.UserID,
		}

		jsonPayload, err := json.Marshal(kafkaPayload)
		if err != nil {
			return errors.Wrap(err, "failed to marshal kafka payload")
		}

		err = s.External.ProduceKafkaMessage(ctx, helpers.GetEnv("KAFKA_TOPIC_PAYMENT_REFUND"), jsonPayload)
		if err != nil {
			return errors.Wrap(err, "failed to produce kafka message")
		}
	}

	return s.OrderRepository.UpdateStatusOrder(ctx, orderID, req.Status)
}

func (s *OrderService) GetOrderList(ctx context.Context) ([]models.Order, error) {
	return s.OrderRepository.GetOrder(ctx)
}

func (s *OrderService) GetOrderDetail(ctx context.Context, orderID int) (*models.Order, error) {
	return s.OrderRepository.GetDetailOrder(ctx, orderID)
}
