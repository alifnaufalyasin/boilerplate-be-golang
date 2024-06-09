package http

import (
	"net/http"
	"time"

	"github.com/alifnaufalyasin/boilerplate-be-golang/src/internal/utils"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// ErrRateLimitExceeded denotes an error raised when rate limit is exceeded
	ErrRateLimitExceeded = echo.NewHTTPError(http.StatusTooManyRequests, "rate limit exceeded")
	// ErrExtractorError denotes an error raised when extractor function is unsuccessful
	ErrExtractorError = echo.NewHTTPError(http.StatusForbidden, "error while extracting identifier")
)

const (
	userClaimKey = "userClaims"
)

type APIServer struct {
	EchoServer     *echo.Echo
	middlewareFunc map[string]echo.MiddlewareFunc
	log            utils.Loggers
}

func Init(e *echo.Echo, secret string, log utils.Loggers) *APIServer {
	middlewareJwt := echojwt.Config{
		BeforeFunc: func(c echo.Context) {
			c.Logger().Info("Token: ", c.Request().Header.Get("Authorization"))
		},
		TokenLookup:    "header:Authorization:Bearer ",
		ParseTokenFunc: utils.GetJWTData,
		SigningKey:     []byte(secret),
		ContextKey:     userClaimKey,
	}

	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}

	// DefaultRateLimiterConfig defines default values for RateLimiterConfig
	var DefaultRateLimiterConfig = middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 5 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return &echo.HTTPError{
				Code:     ErrExtractorError.Code,
				Message:  ErrExtractorError.Message,
				Internal: err,
			}
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return &echo.HTTPError{
				Code:     ErrRateLimitExceeded.Code,
				Message:  ErrRateLimitExceeded.Message,
				Internal: err,
			}
		},
	}

	middlewareFunc := map[string]echo.MiddlewareFunc{
		"cors":        middleware.CORSWithConfig(DefaultCORSConfig),
		"jwt":         echojwt.WithConfig(middlewareJwt),
		"rateLimiter": middleware.RateLimiterWithConfig(DefaultRateLimiterConfig),
	}
	e.Use(middlewareFunc["cors"], middlewareFunc["rateLimiter"])
	e.Logger.Info("menginisialisasikan routes")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!!")
	})

	e.Logger.Info("routes terinisialisasi")

	return &APIServer{
		EchoServer:     e,
		middlewareFunc: middlewareFunc,
		log:            log,
	}
}
