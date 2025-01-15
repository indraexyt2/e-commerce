package services

import (
	"context"
	"e-commerce-product/internal/interfaces"
	"e-commerce-product/internal/models"
	"encoding/json"
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

func (s *CategoryService) UpdateCategory(ctx context.Context, categoryID int, req *models.ProductCategory) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal product")
	}

	newData := map[string]interface{}{}
	err = json.Unmarshal(jsonReq, &newData)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal product")
	}

	err = s.CategoryRepository.UpdateCategory(ctx, categoryID, newData)
	if err != nil {
		return errors.Wrap(err, "failed to update category")
	}

	return nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, categoryID int) error {
	return s.CategoryRepository.DeleteCategory(ctx, categoryID)
}
