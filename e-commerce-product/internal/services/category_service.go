package services

import (
	"context"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/models"
	"github.com/pkg/errors"
)

type CategoryService struct {
	CategoryRepository interfaces.ICategoryRepository
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *models.ProductCategory) (*models.ProductCategory, error) {
	err := s.CategoryRepository.InsertNewCategory(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create category")
	}
	resp := req
	return resp, nil
}
