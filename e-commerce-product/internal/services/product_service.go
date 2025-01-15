package services

import (
	"context"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/models"
	"encoding/json"
	"github.com/pkg/errors"
)

type ProductService struct {
	ProductRepository interfaces.IProductRepository
}

func (s *ProductService) CreateProduct(ctx context.Context, req *models.Product) (*models.Product, error) {
	err := s.ProductRepository.InsertNewProduct(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create product")
	}
	resp := req
	return resp, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, productID int, req *models.Product) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal product")
	}

	newData := map[string]interface{}{}
	err = json.Unmarshal(jsonReq, &newData)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal product")
	}

	err = s.ProductRepository.UpdateProduct(ctx, productID, newData)
	if err != nil {
		return errors.Wrap(err, "failed to update product")
	}

	return nil
}

func (s *ProductService) UpdateProductVariant(ctx context.Context, variantID int, req *models.ProductVariant) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal product")
	}

	newData := map[string]interface{}{}
	err = json.Unmarshal(jsonReq, &newData)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal product")
	}

	err = s.ProductRepository.UpdateProductVariant(ctx, variantID, newData)
	if err != nil {
		return errors.Wrap(err, "failed to update product")
	}

	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, productID int) error {
	return s.ProductRepository.DeleteProduct(ctx, productID)
}
