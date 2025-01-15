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
	"gorm.io/gorm/clause"
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

			err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, product.ID), string(jsonData), time.Hour*24).Err()
			if err != nil {
				helpers.Logger.Warn("Error setting key: ", err)
			}
		}()
	}

	return err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, productID int, newData map[string]interface{}) error {
	err := r.DB.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&models.Product{}).Where("id = ?", productID).
		Updates(newData).Error

	if err != nil {
		return errors.Wrap(err, "failed to update product")
	}

	go func() {
		ctx = context.Background()

		err = r.Redis.Del(ctx, constants.RedisKeyProducts).Err()
		if err != nil {
			helpers.Logger.Warn("Error deleting key: ", err)
		}

		err = r.Redis.Del(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, productID)).Err()
		if err != nil {
			helpers.Logger.Warn("Error deleting key: ", err)
		}
	}()

	return nil
}

func (r *ProductRepository) UpdateProductVariant(ctx context.Context, variantID int, newData map[string]interface{}) error {
	err := r.DB.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&models.ProductVariant{}).Where("id = ?", variantID).
		Updates(newData).Error

	if err != nil {
		return errors.Wrap(err, "failed to update product")
	}

	go func() {
		ctx := context.Background()

		variant := &models.ProductVariant{}
		err = r.DB.Where("id = ?", variantID).Take(variant).Error
		if err != nil {
			helpers.Logger.Warn("Error getting product: ", err)
			return
		}

		err = r.Redis.Del(ctx, constants.RedisKeyProducts).Err()
		if err != nil {
			helpers.Logger.Warn("Error deleting key: ", err)
		}

		err = r.Redis.Del(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, variant.ProductID)).Err()
		if err != nil {
			helpers.Logger.Warn("Error deleting key: ", err)
		}
	}()

	return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, productID int) error {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", productID).Delete(&models.Product{}).Error
		if err != nil {
			return err
		}

		err = tx.Where("product_id = ?", productID).Delete(&models.ProductVariant{}).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err == nil {
		go func() {
			ctx := context.Background()

			err = r.Redis.Del(ctx, constants.RedisKeyProducts).Err()
			if err != nil {
				helpers.Logger.Warn("Error deleting key: ", err)
			}

			err = r.Redis.Del(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, productID)).Err()
			if err != nil {
				helpers.Logger.Warn("Error deleting key: ", err)
			}

		}()
	}

	return err
}

func (r *ProductRepository) GetProducts(ctx context.Context, page int, limit int) ([]*models.Product, error) {
	offset := (page - 1) * limit

	var products []*models.Product
	productStr, err := r.Redis.Get(ctx, constants.RedisKeyProducts).Result()
	if err == nil && productStr != "" {
		var result []*models.Product
		if err := json.Unmarshal([]byte(productStr), &products); err != nil {
			helpers.Logger.Warn("Error unmarshalling products: ", err)
		}
		if page > 0 && limit > 0 {
			for i := offset; i < len(products); i++ {
				if i == offset+limit {
					break
				}
				result = append(result, products[i])
			}
		} else {
			helpers.Logger.Info("Successfully got products from redis")
			return products, nil
		}
		helpers.Logger.Info("Successfully got products from redis")
		return result, nil
	}

	sql := r.DB.WithContext(ctx).Preload("ProductVariants")
	if page > 0 && limit > 0 {
		sql = sql.Offset(offset).Limit(limit)
	}
	err = sql.Find(&products).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get products")
	}

	go func() {
		ctx = context.Background()

		var cacheProducts []*models.Product
		err := sql.Find(&cacheProducts).Error
		if err != nil {
			helpers.Logger.Warn("Error getting products: ", err)
			return
		}

		jsonCacheProduct, err := json.Marshal(cacheProducts)
		if err != nil {
			helpers.Logger.Warn("Error marshalling products: ", err)
			return
		}

		err = r.Redis.Set(ctx, constants.RedisKeyProducts, string(jsonCacheProduct), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("Error setting key: ", err)
			return
		}

		helpers.Logger.Info("Successfully set products to redis")

	}()

	return products, nil
}
