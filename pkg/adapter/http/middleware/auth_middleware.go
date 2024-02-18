package middleware

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/usecase"
	"golang.org/x/xerrors"
)

type Auth interface {
	AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type AuthMiddleware struct {
	AuthUseCase usecase.Auth
}

func NewAuthMiddleware(au usecase.Auth) Auth {
	AuthMiddleware := AuthMiddleware{
		AuthUseCase: au,
	}
	return &AuthMiddleware
}

func (a *AuthMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if a.isSkippedPath(c.Request().URL.Path, c.Request().Referer()) {
			if err := next(c); err != nil {
				return xerrors.Errorf("AuthMiddleware error path: %s: %w", c.Request().URL.Path, err)
			}
			return nil
		}

		cookie, err := c.Cookie("token")
		if err != nil {
			return xerrors.Errorf("AuthMiddleware not extract cookie: %w", err)
		}

		claims, err := a.AuthUseCase.JwtParser(cookie.Value)
		if err != nil {
			a.AuthUseCase.DeleteCookie(c, cookie)
			return fmt.Errorf("failed to parse jwt claims: %w", err)
		}

		var cl = *claims
		uId := int(cl["user_id"].(float64))
		if err := a.AuthUseCase.IdentifyJwtUser(uId); err != nil {
			a.AuthUseCase.DeleteCookie(c, cookie)
			return fmt.Errorf("failed to personal authentication: %w", err)
		}

		if err := next(c); err != nil {
			return xerrors.Errorf("failed to AuthMiddleware err: %w", err)
		}

		return nil
	}
}

func (a *AuthMiddleware) isSkippedPath(reqPath, refPath string) bool {
	skippedPaths := []string{"/healthcheck", "csrf-cookie", "/login", "logout", "playground"}
	for _, path := range skippedPaths {
		if strings.Contains(reqPath, path) || strings.Contains(refPath, path) {
			return true
		}
	}

	return false
}
