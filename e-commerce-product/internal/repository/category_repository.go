package repository

import (
	"context"
	"e-commerce-product/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r *CategoryRepository) InsertNewCategory(ctx context.Context, category *models.ProductCategory) error {
	return r.DB.Create(category).Error
}
