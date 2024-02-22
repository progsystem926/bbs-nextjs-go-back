package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/adapter/http/handler"
	authMiddleware "github.com/progsystem926/bbs-nextjs-go-back/pkg/adapter/http/middleware"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/validator"
	"golang.org/x/xerrors"
)

type Router interface {
	InitRouting(*config.Config) (*echo.Echo, error)
}

type InitRouter struct {
	Ch handler.Csrf
	Lh handler.Login
	Gh handler.Graph
	Sh handler.SignUp
	Ph http.HandlerFunc
	Am authMiddleware.Auth
}

func NewInitRouter(ch handler.Csrf, lh handler.Login, gh handler.Graph, sh handler.SignUp, ph http.HandlerFunc, am authMiddleware.Auth) Router {
	InitRouter := InitRouter{ch, lh, gh, sh, ph, am}
	return &InitRouter
}

func (i *InitRouter) InitRouting(cfg *config.Config) (*echo.Echo, error) {
	e := echo.New()

	cookieDomain := ""
	if cfg.Env == "prd" {
		cookieDomain = "." + cfg.AppDomain
	}

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{cfg.FrontURL},
			AllowCredentials: true,
			AllowHeaders: []string{
				echo.HeaderContentType,
				echo.HeaderXCSRFToken,
			},
		}),
		middleware.CSRFWithConfig(middleware.CSRFConfig{
			CookiePath:     "/",
			CookieSecure:   true,
			CookieDomain:   cookieDomain,
			CookieSameSite: http.SameSiteNoneMode,
			Skipper: func(c echo.Context) bool {
				if strings.Contains(c.Request().URL.Path, "/healthcheck") {
					return true
				}
				if strings.Contains(c.Request().URL.Path, "/playground") {
					return true
				}
				if strings.Contains(c.Request().URL.Path, "/query") {
					return true
				}
				return false
			},
		}),
		i.Am.AuthMiddleware,
	)

	e.Validator = validator.NewValidator()

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "New deployment test")
	})
	e.GET("/csrf-cookie", i.Ch.CsrfHandler())
	e.POST("/login", i.Lh.LoginHandler())
	e.GET("/logout", i.Lh.LogoutHandler())
	e.POST("/signup", i.Sh.SignUpHandler())
	e.POST("/query", i.Gh.QueryHandler())
	e.GET("/playground", func(c echo.Context) error {
		i.Ph.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		return nil, xerrors.Errorf("fail to start port:%s %w", cfg.Port, err)
	}

	return e, nil
}

func customHTTPErrorHandler(err error, c echo.Context) {
	sentry.CaptureException(fmt.Errorf("handler err: %w", err))

	c.Logger().Error(err)

	if err := c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	}); err != nil {
		c.Logger().Error(err)
	}
}
