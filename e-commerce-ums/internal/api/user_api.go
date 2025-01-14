package api

import (
	"e-commerce-ums/helpers"
	"e-commerce-ums/internal/interfaces"
	"e-commerce-ums/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserAPI struct {
	UserService interfaces.IUserService
}

func (api *UserAPI) RegisterUser(e echo.Context) error {
	var (
		req = &models.User{}
		log = helpers.Logger
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	resp, err := api.UserService.Register(e.Request().Context(), req, "user")
	if err != nil {
		log.Error("Error registering user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error registering user. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "User registered successfully", resp)
}

func (api *UserAPI) RegisterAdmin(e echo.Context) error {
	var (
		req = &models.User{}
		log = helpers.Logger
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check try again", nil)
	}

	resp, err := api.UserService.Register(e.Request().Context(), req, "admin")
	if err != nil {
		log.Error("Error registering user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error registering user. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "User registered successfully", resp)
}

func (api *UserAPI) LoginUser(e echo.Context) error {
	var (
		req = &models.LoginRequest{}
		log = helpers.Logger
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	resp, err := api.UserService.Login(e.Request().Context(), req, "user")
	if err != nil {
		log.Error("Error logging in user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error logging in user. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "User logged in successfully", resp)
}

func (api *UserAPI) LoginAdmin(e echo.Context) error {
	var (
		req = &models.LoginRequest{}
		log = helpers.Logger
	)

	if err := e.Bind(req); err != nil {
		log.Error("Error binding request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please try again", nil)
	}

	resp, err := api.UserService.Login(e.Request().Context(), req, "admin")
	if err != nil {
		log.Error("Error logging in user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error logging in user. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "User logged in successfully", resp)
}

func (api *UserAPI) GetProfile(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	token := e.Get("token")
	tokenClaims, ok := token.(*helpers.ClaimToken)
	if !ok {
		log.Error("Error getting token claims")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	resp, err := api.UserService.GetProfile(e.Request().Context(), tokenClaims.Username)
	if err != nil {
		log.Error("Error getting user profile: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error getting user profile. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "User profile retrieved successfully", resp)
}

func (api *UserAPI) RefreshToken(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	refreshToken := e.Request().Header.Get("Authorization")
	claim := e.Get("token")
	claimsToken, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("Error getting token claims")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Server error", nil)
	}

	resp, err := api.UserService.RefreshToken(e.Request().Context(), refreshToken, claimsToken)
	if err != nil {
		log.Error("failed to refresh token: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error refreshing token. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "Token refreshed successfully", resp)
}

func (api *UserAPI) Logout(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	token := e.Request().Header.Get("Authorization")
	err := api.UserService.Logout(e.Request().Context(), token)
	if err != nil {
		log.Error("failed to logout: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error logging out. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, "User logged out successfully", nil)
}
