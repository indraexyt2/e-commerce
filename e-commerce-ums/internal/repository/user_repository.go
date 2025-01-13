package repository

import (
	"context"
	"e-commerce-ums/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string, role string) (*models.User, error) {
	var (
		user *models.User
		err  error
	)
	err = r.DB.WithContext(ctx).Where("username = ?", username).Where("role = ?", role).First(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.WithContext(ctx).Create(session).Error
}
