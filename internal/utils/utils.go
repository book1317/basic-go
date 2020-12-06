package utils

import (
	"net/http"

	"github.com/labstack/echo"

	"academy/internal/model"
)

func JSONResponse(c echo.Context, data interface{}, err error) error {
	code, message := "200", "OK"
	if err != nil {
		code, message = "500", err.Error()
	}
	res := model.Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}
