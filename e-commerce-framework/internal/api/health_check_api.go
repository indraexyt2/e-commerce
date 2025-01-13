package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckAPI struct {
}

func (api *HealthCheckAPI) HealthCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}
