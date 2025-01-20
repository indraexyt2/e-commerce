package services

import (
	"context"
	"e-commerce-payment/external"
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

func (s *PaymentService) InitiatePayment(ctx context.Context, req *models.PaymentInitiatePayload) error {
	paymentMethod, err := s.PaymentRepository.GetPaymentMethod(ctx, req.UserID, "e-commerce")
	if err != nil {
		return errors.Wrap(err, "failed to get payment method")
	}

	trxReq := &external.PaymentTransactionRequest{
		WalletID:        paymentMethod.SourceID,
		Amount:          req.TotalPrice,
		Reference:       helpers.GenerateReference(),
		TransactionType: "DEBIT",
	}

	resp, err := s.External.WalletTransaction(ctx, trxReq)
	if err != nil {
		return errors.Wrap(err, "failed to proceed to wallet transaction")
	}
	helpers.Logger.WithField("balance", resp.Data.Balance).Info("Succeed to payment")

	paymentTrx := &models.PaymentTransaction{
		UserID:           req.UserID,
		OrderID:          req.OrderID,
		TotalPrice:       req.TotalPrice,
		PaymentMethodID:  paymentMethod.ID,
		Status:           "SUCCESS",
		PaymentReference: trxReq.Reference,
	}
	err = s.PaymentRepository.InsertNewPaymentTransaction(ctx, paymentTrx)
	if err != nil {
		return errors.Wrap(err, "failed to insert new payment transaction")
	}

	_, err = s.External.OrderCallback(ctx, req.OrderID, "SUCCESS")
	if err != nil {
		return errors.Wrap(err, "failed to callback order")
	}

	return nil
}

func (s *PaymentService) RefundPayment(ctx context.Context, req *models.RefundPayload) error {
	paymentDetail, err := s.PaymentRepository.GetPaymentByOrderID(ctx, req.OrderID)
	if err != nil {
		return errors.Wrap(err, "failed to get payment detail")
	}

	paymentMethod, err := s.PaymentRepository.GetPaymentMethodById(ctx, paymentDetail.PaymentMethodID)
	if err != nil {
		return errors.Wrap(err, "failed to get payment method")
	}

	trxReq := &external.PaymentTransactionRequest{
		WalletID:        paymentMethod.SourceID,
		Amount:          paymentDetail.TotalPrice,
		Reference:       "REFUND" + "-" + paymentDetail.PaymentReference,
		TransactionType: "CREDIT",
	}

	resp, err := s.External.WalletTransaction(ctx, trxReq)
	if err != nil {
		return errors.Wrap(err, "failed to proceed to wallet transaction")
	}
	helpers.Logger.WithField("balance", resp.Data.Balance).Info("Succeed to refund")

	refund := &models.PaymentRefund{
		AdminID:          req.AdminID,
		OrderID:          req.OrderID,
		Status:           "SUCCESS",
		PaymentReference: trxReq.Reference,
	}
	err = s.PaymentRepository.InsertNewPaymentRefund(ctx, refund)
	if err != nil {
		return errors.Wrap(err, "failed to refund transaction")
	}

	return nil
}
