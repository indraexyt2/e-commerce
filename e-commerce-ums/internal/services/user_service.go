package services

import (
	"context"
	"e-commerce-ums/helpers"
	"e-commerce-ums/internal/interfaces"
	"e-commerce-ums/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func (s *UserService) Register(ctx context.Context, req *models.User, role string) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashPassword)
	req.Role = role

	err = s.UserRepository.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := req
	resp.Password = ""
	return resp, nil
}

func (s *UserService) Login(ctx context.Context, request *models.LoginRequest, role string) (*models.LoginResponse, error) {
	var (
		resp = &models.LoginResponse{}
		now  = time.Now()
	)

	userDetail, err := s.UserRepository.GetUserByUsername(ctx, request.Username, role)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get user detail")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password)); err != nil {
		return resp, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, userDetail.Email, "token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, userDetail.Email, "refresh_token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepository.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return resp, errors.Wrap(err, "failed to insert new user session")
	}

	resp.UserID = userDetail.ID
	resp.Username = userDetail.Username
	resp.FullName = userDetail.FullName
	resp.Email = userDetail.Email
	resp.Token = token
	resp.RefreshToken = refreshToken
	return resp, nil
}

func (s *UserService) GetProfile(ctx context.Context, username string) (*models.User, error) {
	resp, err := s.UserRepository.GetUserByUsername(ctx, username, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user detail")
	}

	resp.Password = ""
	return resp, nil
}

func (s *UserService) RefreshToken(ctx context.Context, refreshToken string, claimToken *helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	resp := models.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, claimToken.UserID, claimToken.Username, claimToken.FullName, claimToken.Email, "token", time.Now())
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	err = s.UserRepository.UpdateTokenByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return resp, errors.Wrap(err, "failed to update token")
	}

	resp.Token = token
	return resp, nil
}

func (s *UserService) Logout(ctx context.Context, token string) error {
	return s.UserRepository.DeleteUserSession(ctx, token)
}
