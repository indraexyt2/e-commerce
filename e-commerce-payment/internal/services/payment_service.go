package services

import (
	"context"
	"e-commerce-payment/helpers"
	"e-commerce-payment/internal/interfaces"
	"e-commerce-payment/internal/models"
	"github.com/pkg/errors"
)

type PaymentService struct {
	PaymentRepository interfaces.IPaymentRepository
	External          interfaces.IExternal
}

func (s *PaymentService) PaymentMethodLink(ctx context.Context, req *models.PaymentMethodLink) error {
	resp, err := s.External.PaymentLink(ctx, req)
	if err != nil {
		return errors.Wrap(err, "failed to link payment")
	}
	helpers.Logger.WithField("OTP", resp.Data.OTP).Info("Link response is success, need otp confirmation")
	return nil
}

func (s *PaymentService) PaymentMethodLinkConfirm(ctx context.Context, userID int, req *models.PaymentMethodLinkConfirm) error {
	_, err := s.External.PaymentLinkConfirmation(ctx, req.SourceID, req.OTP)
	if err != nil {
		return errors.Wrap(err, "failed to link confirmation")
	}

	paymentMethod := &models.PaymentMethod{
		UserID:     userID,
		SourceID:   req.SourceID,
		SourceName: "e-commerce",
	}
	return s.PaymentRepository.InsertNewPaymentMethod(ctx, paymentMethod)
}

func (s *PaymentService) PaymentMethodUnlink(ctx context.Context, userID int, req *models.PaymentMethodLink) error {
	_, err := s.External.PaymentUnlink(ctx, req)
	if err != nil {
		return errors.Wrap(err, "failed to unlink payment")
	}

	return s.PaymentRepository.DeletePaymentMethod(ctx, req.SourceID, userID, "e-commerce")
}
