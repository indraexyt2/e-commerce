package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"time"
)

type Product struct {
	ID              int              `json:"id"`
	Name            string           `json:"name" gorm:"column:name;type:varchar(255)" validate:"required"`
	Description     string           `json:"description" gorm:"column:description;type:text" validate:"required"`
	Price           float64          `json:"price" gorm:"column:price;type:decimal(10,2)" validate:"required"`
	Categories      pq.Int64Array    `json:"categories,omitempty" gorm:"column:categories;type:int[]" validate:"required"`
	CreatedAt       time.Time        `json:"-"`
	UpdatedAt       time.Time        `json:"-"`
	ProductVariants []ProductVariant `json:"variants,omitempty" gorm:"-"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Validate() error {
	v := validator.New()
	return v.Struct(p)
}

type ProductCategory struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255)" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (c *ProductCategory) TableName() string {
	return "product_categories"
}

func (c *ProductCategory) Validate() error {
	v := validator.New()
	return v.Struct(c)
}

type ProductVariant struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id" gorm:"column:product_id;type:int" validate:"required"`
	Color     string    `json:"color" gorm:"column:color;type:varchar(50)" validate:"required"`
	Size      string    `json:"size" gorm:"column:size;type:varchar(50)" validate:"required"`
	Quantity  int       `json:"quantity" gorm:"column:quantity;type:int"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (p *ProductVariant) TableName() string {
	return "product_variants"
}
