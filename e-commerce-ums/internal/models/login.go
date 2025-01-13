package models

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (I *LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type LoginResponse struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
