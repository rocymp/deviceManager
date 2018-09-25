package common

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResultJson struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func JSONE(c echo.Context, msg string, code int, data interface{}) error {
	res := &ResultJson{
		Code:    code,
		Message: msg,
		Data:    data,
	}

	return c.JSON(http.StatusOK, res)
}
