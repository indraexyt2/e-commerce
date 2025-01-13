package services

import (
	"context"
	"e-commerce-ums/internal/interfaces"
	"e-commerce-ums/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func (s *UserService) RegisterUser(ctx context.Context, req *models.User) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashPassword)
	req.Role = "user"

	err = s.UserRepository.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := req
	resp.Password = ""
	return resp, nil
}

func (s *UserService) RegisterAdmin(ctx context.Context, req *models.User) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashPassword)
	req.Role = "admin"

	err = s.UserRepository.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := req
	resp.Password = ""
	return resp, nil
}
