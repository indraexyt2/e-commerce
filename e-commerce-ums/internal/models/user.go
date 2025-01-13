package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID          int       `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Username    string    `json:"username" gorm:"column:username;unique_index;type:varchar(20)" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;unique_index;type:varchar(100)" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;unique_index;type:varchar(15)" validate:"required"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(100)"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255)" validate:"required"`
	Role        string    `json:"role" gorm:"column:role;type:user_role" validate:"required"`
	CreatedAt   time.Time `json:"-" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"-" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (*User) TableName() string {
	return "users"
}

func (I *User) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type UserSession struct {
	ID                  int `json:"id" gorm:"column:id;primary_key;auto_increment"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"column:user_id;type:int" validate:"required"`
	Token               string    `json:"token" gorm:"column:token;type:varchar(500)" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"column:refresh_token;type:varchar(500)" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (I *UserSession) TableName() string {
	return "user_sessions"
}

func (I *UserSession) Validate() error {
	v := validator.New()
	return v.Struct(I)
}
