package api

import (
	"e-commerce-order/external"
	"e-commerce-order/helpers"
	"e-commerce-order/internal/interfaces"
	"e-commerce-order/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
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
