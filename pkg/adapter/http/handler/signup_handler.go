package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/usecase"
	"golang.org/x/xerrors"
)

type SignUp interface {
	SignUpHandler() echo.HandlerFunc
}

type SignUpHandler struct {
	UserUseCase usecase.User
}

func NewSignUpHandler(uu usecase.User) SignUp {
	SignUpHandler := SignUpHandler{
		UserUseCase: uu,
	}
	return &SignUpHandler
}

func (s *SignUpHandler) SignUpHandler() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var fv = &usecase.SignUpFormValue{
			Name:     c.FormValue("name"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		if err = c.Validate(fv); err != nil {
			return xerrors.Errorf("SignUp validate err: %w", err)
		}

		userId, err := s.UserUseCase.SignUp(c, fv)
		if err != nil {
			return fmt.Errorf("SignUp failed err: %w", err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"id": userId,
		})
	}
}
