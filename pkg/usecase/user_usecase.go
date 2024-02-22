package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/repository"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	"golang.org/x/xerrors"
)

type User interface {
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	SignUp(c echo.Context, fv *SignUpFormValue) (int, error)
}

type UserUseCase struct {
	userRepo repository.User
	config   *config.Config
}

type SignUpFormValue struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewUserUseCase(userRepo repository.User, config *config.Config) User {
	UserUseCase := UserUseCase{
		userRepo: userRepo,
		config:   config,
	}
	return &UserUseCase
}

func (u *UserUseCase) GetUserById(id int) (*model.User, error) {
	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf("useCase GetUserById() err: %w", err)
	}

	return user, nil
}

func (u *UserUseCase) GetUserByEmail(email string) (*model.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("useCase GetUserByEmail() err: %w", err)
	}

	return user, nil
}

func (u *UserUseCase) SignUp(c echo.Context, fv *SignUpFormValue) (int, error) {
	encEmail, err := u.userRepo.Encrypt(fv.Email)
	if err != nil {
		return 0, fmt.Errorf("SignUp failed at Encrypt email err: %w", err)
	}

	encPass, err := u.userRepo.Encrypt(fv.Password)
	if err != nil {
		return 0, fmt.Errorf("SignUp failed at Encrypt password err: %w", err)
	}

	user := model.User{
		Name:     fv.Name,
		Email:    encEmail,
		Password: encPass,
	}

	_, err = u.userRepo.CreateUser(&user)
	if err != nil {
		return 0, fmt.Errorf("SignUp failed at CreateUser err: %w", err)
	}

	created, err := u.userRepo.GetUserByEmail(encEmail)
	if err != nil {
		return 0, fmt.Errorf("SignUp failed at GetUserByEmail err: %w", err)
	}

	claims := &jwtCustomClaims{
		created.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(u.config.JwtSecret))
	if err != nil {
		return 0, xerrors.Errorf("login failed at NewWithClaims err %w", err)
	}

	time.Local = time.FixedZone("Local", 9*60*60)
	jst, err := time.LoadLocation("Local")
	if err != nil {
		return 0, xerrors.Errorf("login failed at LoadLocation err %w", err)
	}
	nowJST := time.Now().In(jst)

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = nowJST.Add(1 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return created.ID, nil
}
