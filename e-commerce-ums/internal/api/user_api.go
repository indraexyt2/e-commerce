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

	resp, err := api.UserService.RegisterUser(e.Request().Context(), req)
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
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, "Invalid input. Please check your data and try again", nil)
	}

	resp, err := api.UserService.RegisterAdmin(e.Request().Context(), req)
	if err != nil {
		log.Error("Error registering user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Error registering user. Please try again", nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusCreated, "User registered successfully", resp)
}
