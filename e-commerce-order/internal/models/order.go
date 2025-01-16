package models

import "github.com/go-playground/validator/v10"

type Order struct {
	ID         int         `json:"id"`
	UserID     int         `json:"user_id"`
	TotalPrice float64     `json:"total_price" gorm:"column:total_price;type:decimal(10,2)" validate:"required"`
	Status     string      `json:"status" gorm:"column:status;type:varchar(10)" validate:"required"`
	CreatedAt  string      `json:"-"`
	UpdatedAt  string      `json:"-"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) Validate() error {
	v := validator.New()
	return v.Struct(o)
}

type OrderItem struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id" validate:"required"`
	VariantID int     `json:"variant_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price     float64 `json:"price" gorm:"column:price;type:decimal(10,2)" validate:"required"`
	CreatedAt string  `json:"-"`
	UpdatedAt string  `json:"-"`
}

func (oi *OrderItem) TableName() string {
	return "order_items"
}

func (oi *OrderItem) Validate() error {
	v := validator.New()
	return v.Struct(oi)
}
