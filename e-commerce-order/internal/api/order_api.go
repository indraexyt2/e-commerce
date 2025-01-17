package api

import (
	"e-commerce-order/external"
	"e-commerce-order/helpers"
	"e-commerce-order/internal/interfaces"
	"e-commerce-order/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type OrderAPI struct {
	OrderService interfaces.IOrderService
}

func (api *OrderAPI) CreateOrder(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	req := models.Order{}
	if err := e.Bind(&req); err != nil {
		log.Error("failed to bind request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid data", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid data", nil)
	}

	profileCtx := e.Get("profile")
	profile, ok := profileCtx.(*external.Profile)
	if !ok {
		log.Error("failed to get profile from context")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	resp, err := api.OrderService.CreateOrder(e.Request().Context(), profile, &req)
	if err != nil {
		log.Error("failed to create order: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Success", resp)
}

func (api *OrderAPI) UpdateOrderStatus(e echo.Context) error {
	var (
		log        = helpers.Logger
		req        = &models.OrderStatusRequest{}
		orderIDStr = e.Param("id")
		profileCtx = e.Get("profile")
	)

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		log.Error("failed to convert order id to int: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid data", nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("failed to bind request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid data", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid data", nil)
	}

	profile, ok := profileCtx.(*external.Profile)
	if !ok {
		log.Error("failed to get profile from context")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	err = api.OrderService.UpdateOrderStatus(e.Request().Context(), orderID, profile, req)
	if err != nil {
		log.Error("failed to update order status: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Success", nil)
}

func (api *OrderAPI) GetOrderDetail(e echo.Context) error {
	var (
		log        = helpers.Logger
		orderIDStr = e.Param("id")
	)

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		log.Error("failed to convert order id to int: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid data", nil)
	}

	resp, err := api.OrderService.GetOrderDetail(e.Request().Context(), orderID)
	if err != nil {
		log.Error("failed to get order detail: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Success", resp)
}

func (api *OrderAPI) GetOrderList(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	resp, err := api.OrderService.GetOrderList(e.Request().Context())
	if err != nil {
		log.Error("failed to get order list: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Success", resp)
}
