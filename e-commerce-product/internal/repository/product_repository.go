package repository

import (
	"context"
	"e-commerce-product/constants"
	"e-commerce-product/helpers"
	"e-commerce-product/internal/models"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type ProductRepository struct {
	DB    *gorm.DB
	Redis *redis.ClusterClient
}

func (r *ProductRepository) InsertNewProduct(ctx context.Context, product *models.Product) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(product).Error
		if err != nil {
			return err
		}

		for i, variant := range product.ProductVariants {
			variant.ProductID = product.ID
			err := tx.Create(&variant).Error
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to create product variant %v", variant))
			}

			product.ProductVariants[i].ID = variant.ID
			product.ProductVariants[i].ProductID = product.ID
		}
		return nil
	})

	if err == nil {
		go func() {
			ctx := context.Background()
			jsonData, err := json.Marshal(product)
			if err != nil {
				helpers.Logger.Warn("Error marshalling product: ", err)
				return
			}

			err = r.Redis.Del(ctx, constants.RedisKeyProducts).Err()
			if err != nil {
				helpers.Logger.Warn("Error deleting key: ", err)
			}
			err = r.Redis.Del(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, product.ID)).Err()
			if err != nil {
				helpers.Logger.Warn("Error deleting key: ", err)
			}
			err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, product.ID), string(jsonData), time.Hour*24).Err()
			if err != nil {
				helpers.Logger.Warn("Error setting key: ", err)
			}
		}()
	}

	return err
}
