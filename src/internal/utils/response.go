package utils

import (
	"github.com/labstack/echo/v4"
)

type JSONResponseData struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type JSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Response(c echo.Context, res JSONResponse) error {
	return c.JSON(int(res.Code), res)
}

func ResponseData(c echo.Context, res JSONResponseData) error {
	return c.JSON(int(res.Code), res)
}
