package repository

import (
	"context"
	"e-commerce-product/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r *CategoryRepository) InsertNewCategory(ctx context.Context, category *models.ProductCategory) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, categoryID int, newData map[string]interface{}) error {
	err := r.DB.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&models.ProductCategory{}).Where("id = ?", categoryID).
		Updates(newData).Error

	if err != nil {
		return errors.Wrap(err, "failed to update product")
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, categoryID int) error {
	return r.DB.WithContext(ctx).Delete(&models.ProductCategory{}, "id = ?", categoryID).Error
}
