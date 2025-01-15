package api

import (
	"e-commerce-product/helpers"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func (api *CategoryAPI) UpdateProductCategory(e echo.Context) error {
	var (
		req           = &models.ProductCategory{}
		log           = helpers.Logger
		categoryIDStr = e.Param("id")
	)

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil || categoryID == 0 {
		log.Error("Error parsing category ID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	err = api.CategoryService.UpdateCategory(e.Request().Context(), categoryID, req)
	if err != nil {
		log.Error("Error creating category: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error updating category. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "Category updated successfully", nil)
}

func (api *CategoryAPI) DeleteCategory(e echo.Context) error {
	var (
		log           = helpers.Logger
		categoryIDStr = e.Param("id")
	)

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil || categoryID == 0 {
		log.Error("Error parsing category ID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	err = api.CategoryService.DeleteCategory(e.Request().Context(), categoryID)
	if err != nil {
		log.Error("Error deleting category: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error deleting category. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Category deleted successfully", nil)
}

func (api *CategoryAPI) GetCategories(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	resp, err := api.CategoryService.GetCategories(e.Request().Context())
	if err != nil {
		log.Error("Error getting categories: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error getting categories. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Categories retrieved successfully", resp)
}
