package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/alifnaufalyasin/boilerplate-be-golang/src/cmd/service/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Username string `json:"username"`
	UserId   string `json:"id"`
	jwt.StandardClaims
}

var (
	ExpiredHour int = 24
)

func (j JWTCustomClaims) Valid() error {
	return nil
}

func GenerateToken(secret, nama, username, email, id string, ExpiredHour int) (string, error) {
	timeNow := time.Now()
	JWTStandartClaims := jwt.StandardClaims{
		ExpiresAt: timeNow.Add(time.Hour * 24).Unix(),
		IssuedAt:  timeNow.Unix(),
		Issuer:    "aliven-server",
	}

	// Set custom claims
	claims := &JWTCustomClaims{
		Nama:           nama,
		Username:       username,
		Email:          email,
		UserId:         id,
		StandardClaims: JWTStandartClaims,
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetJWTData(c echo.Context, auth string) (interface{}, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, errors.New("unexpected error on getting config")
	}
	finalToken, err := jwt.ParseWithClaims(auth, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return nil, errors.New("unexpected error on getting jwt data")
	}

	claims := finalToken.Claims
	return claims, nil
}
