package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type PaymentInitiatePayload struct {
	UserID     int     `json:"user_id"`
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}

type RefundPayload struct {
	OrderID int `json:"order_id"`
	AdminID int `json:"admin_id"`
}

type PaymentTransaction struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	OrderID          int       `json:"order_id" validate:"required"`
	TotalPrice       float64   `json:"total_price" gorm:"column:total_price;type:decimal(10,2)" validate:"required"`
	PaymentMethodID  int       `json:"payment_method_id"`
	Status           string    `json:"status" gorm:"column:status;type:varchar(10)"`
	PaymentReference string    `json:"payment_reference" gorm:"column:payment_reference;type:varchar(255)"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

func (pt *PaymentTransaction) TableName() string {
	return "payment_transactions"
}

func (pt *PaymentTransaction) Validate() error {
	v := validator.New()
	return v.Struct(pt)
}

type PaymentMethod struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	SourceID   int       `json:"source_id"`
	SourceName string    `json:"source_name" gorm:"column:source_name;type:varchar(50)"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

func (pm *PaymentMethod) TableName() string {
	return "payment_methods"
}

func (pm *PaymentMethod) Validate() error {
	v := validator.New()
	return v.Struct(pm)
}

type PaymentRefund struct {
	ID               int       `json:"id"`
	AdminID          int       `json:"admin_id"`
	OrderID          int       `json:"order_id" validate:"required"`
	Status           string    `json:"status" gorm:"column:status;type:varchar(10)"`
	PaymentReference string    `json:"payment_reference" gorm:"column:payment_reference;type:varchar(255)"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

func (pr *PaymentRefund) TableName() string {
	return "payment_refunds"
}

func (pr *PaymentRefund) Validate() error {
	v := validator.New()
	return v.Struct(pr)
}

type PaymentMethodLink struct {
	SourceID int `json:"source_id" validate:"required"`
}

func (pml *PaymentMethodLink) Validate() error {
	v := validator.New()
	return v.Struct(pml)
}

type PaymentMethodLinkConfirm struct {
	OTP      string `json:"otp" validate:"required"`
	SourceID int    `json:"source_id" validate:"required"`
}

func (pmc *PaymentMethodLinkConfirm) Validate() error {
	v := validator.New()
	return v.Struct(pmc)
}
