package helpers

import "github.com/labstack/echo/v4"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponseHTTP(e echo.Context, code int, message string, data interface{}) {
	resp := Response{
		Message: message,
		Data:    data,
	}
	e.JSON(code, resp)
}
