package api

import (
	"context"
	"e-commerce-payment/external"
	"e-commerce-payment/helpers"
	"e-commerce-payment/internal/interfaces"
	"e-commerce-payment/internal/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type PaymentAPI struct {
	PaymentService interfaces.IPaymentService
}

func (api *PaymentAPI) PaymentMethodLink(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.PaymentMethodLink{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	err := api.PaymentService.PaymentMethodLink(e.Request().Context(), req)
	if err != nil {
		log.Error("Error linking payment method: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Failed to link payment method. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Payment method linked successfully", nil)
}

func (api *PaymentAPI) PaymentMethodLinkConfirm(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.PaymentMethodLinkConfirm{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	profileCtx := e.Get("profile")
	profile, ok := profileCtx.(*external.Profile)
	if !ok {
		log.Error("Error getting profile from context")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Failed to get profile. Please try again", nil)
	}

	err := api.PaymentService.PaymentMethodLinkConfirm(e.Request().Context(), profile.Data.UserID, req)
	if err != nil {
		log.Error("Error confirming payment method: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Failed to confirm payment method. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Payment method confirmed successfully", nil)
}

func (api *PaymentAPI) PaymentMethodUnlink(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.PaymentMethodLink{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	profileCtx := e.Get("profile")
	profile, ok := profileCtx.(*external.Profile)
	if !ok {
		log.Error("Error getting profile from context")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Failed to get profile. Please try again", nil)
	}

	err := api.PaymentService.PaymentMethodUnlink(e.Request().Context(), profile.Data.UserID, req)
	if err != nil {
		log.Error("Error unlinking payment method: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Failed to confirm payment method. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Payment method unlinked successfully", nil)
}

func (api *PaymentAPI) InitiatePayment(kafkaPayload []byte) error {
	var (
		log = helpers.Logger
		req = &models.PaymentInitiatePayload{}
	)

	if err := json.Unmarshal(kafkaPayload, &req); err != nil {
		log.Error("failed to unmarshal kafka payload: ", err)
		return errors.Wrap(err, "failed to unmarshal kafka payload")
	}

	return api.PaymentService.InitiatePayment(context.Background(), req)
}

func (api *PaymentAPI) RefundPayment(kafkaPayload []byte) error {
	var (
		log = helpers.Logger
		req = &models.RefundPayload{}
	)

	if err := json.Unmarshal(kafkaPayload, &req); err != nil {
		log.Error("failed to unmarshal kafka payload: ", err)
		return errors.Wrap(err, "failed to unmarshal kafka payload")
	}

	return api.PaymentService.RefundPayment(context.Background(), req)
}
