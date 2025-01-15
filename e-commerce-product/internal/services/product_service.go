package services

import (
	"context"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/models"
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
