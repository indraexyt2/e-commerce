package api

import (
	"e-commerce-product/helpers"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryAPI struct {
	CategoryService interfaces.ICategoryService
}

func (api *CategoryAPI) CreateCategory(e echo.Context) error {
	var (
		req = &models.ProductCategory{}
		log = helpers.Logger
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	resp, err := api.CategoryService.CreateCategory(e.Request().Context(), req)
	if err != nil {
		log.Error("Error creating category: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error creating category. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "Category created successfully", resp)
}
