package api

import (
	"e-commerce-product/helpers"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductService interfaces.IProductService
}

func (api *ProductAPI) CreateProduct(e echo.Context) error {
	var (
		req = &models.Product{}
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

	resp, err := api.ProductService.CreateProduct(e.Request().Context(), req)
	if err != nil {
		log.Error("Error creating product: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error creating product. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "Product created successfully", resp)
}

func (api *ProductAPI) UpdateProduct(e echo.Context) error {
	var (
		req          = &models.Product{}
		log          = helpers.Logger
		productIDstr = e.Param("id")
	)

	productID, err := strconv.Atoi(productIDstr)
	if err != nil || productID == 0 {
		log.Error("Error parsing product ID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	err = api.ProductService.UpdateProduct(e.Request().Context(), productID, req)
	if err != nil {
		log.Error("Error updating product: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error updating product. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "Product updated successfully", nil)
}

func (api *ProductAPI) UpdateProductVariant(e echo.Context) error {
	var (
		req          = &models.ProductVariant{}
		log          = helpers.Logger
		variantIDStr = e.Param("id")
	)

	variantID, err := strconv.Atoi(variantIDStr)
	if err != nil || variantID == 0 {
		log.Error("Error parsing product ID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	err = api.ProductService.UpdateProductVariant(e.Request().Context(), variantID, req)
	if err != nil {
		log.Error("Error updating product variant: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error updating product variant. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "Product variant updated successfully", nil)
}

func (api *ProductAPI) DeleteProduct(e echo.Context) error {
	var (
		log          = helpers.Logger
		productIDStr = e.Param("id")
	)

	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID == 0 {
		log.Error("Error parsing product ID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	err = api.ProductService.DeleteProduct(e.Request().Context(), productID)
	if err != nil {
		log.Error("Error deleting product: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error deleting product. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Product deleted successfully", nil)
}

func (api *ProductAPI) GetProducts(e echo.Context) error {
	var (
		log      = helpers.Logger
		pageStr  = e.QueryParam("page")
		limitStr = e.QueryParam("limit")
	)

	page, err := strconv.Atoi(pageStr)
	if err != nil && pageStr != "" {
		log.Error("Error parsing page: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		log.Error("Error parsing limit: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	resp, err := api.ProductService.GetProducts(e.Request().Context(), page, limit)
	if err != nil {
		log.Error("Error getting products: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error getting products. Please try again", nil)
	}

	fmt.Println("response", resp)

	return helpers.SendResponseHTTP(e, http.StatusOK, "Products retrieved successfully", resp)

}

func (api *ProductAPI) GetProductDetail(e echo.Context) error {
	var (
		log          = helpers.Logger
		productIDStr = e.Param("id")
	)

	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID == 0 {
		log.Error("Error parsing product ID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	resp, err := api.ProductService.GetProductDetail(e.Request().Context(), productID)
	if err != nil {
		log.Error("Error getting product detail: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error getting product detail. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Product detail retrieved successfully", resp)
}
