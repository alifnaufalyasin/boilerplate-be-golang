package http

import (
	"net/http"

	"github.com/alifnaufalyasin/boilerplate-be-golang/src/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (a *APIServer) User() {
	a.EchoServer.POST("/users", a.CreateUser)
	a.EchoServer.GET("/users/me", a.GetDetailUser, a.middlewareFunc["jwt"])
}

func (a *APIServer) CreateUser(c echo.Context) error {
	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "berhasil membuat akun",
	})
}

func (a *APIServer) GetDetailUser(c echo.Context) error {

	claims := c.Get(userClaimKey).(jwt.MapClaims)

	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusOK,
		Data:    claims,
		Message: "berhasil memperbarui data",
	})
}
