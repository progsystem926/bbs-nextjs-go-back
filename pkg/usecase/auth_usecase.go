package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/repository"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	"golang.org/x/xerrors"
)

type Auth interface {
	Login(c echo.Context, fv *LoginFormValue) (int, error)
	JwtParser(auth string) (*jwt.MapClaims, error)
	IdentifyJwtUser(id int) error
	DeleteCookie(c echo.Context, ck *http.Cookie)
}

type AuthUseCase struct {
	config   *config.Config
	userRepo repository.User
}

type LoginFormValue struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type jwtCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func NewAuthUseCase(userRepo repository.User, config *config.Config) Auth {
	AuthUseCase := AuthUseCase{
		config:   config,
		userRepo: userRepo,
	}
	return &AuthUseCase
}

func (a *AuthUseCase) JwtParser(auth string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, xerrors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.config.JwtSecret), nil
	})

	if err != nil {
		return nil, xerrors.Errorf("JwtParser failed: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, xerrors.Errorf("invalid claims: %v", claims)
	}

	return &claims, nil
}

func (a *AuthUseCase) Login(c echo.Context, fv *LoginFormValue) (int, error) {
	encEmail, err := a.userRepo.Encrypt(fv.Email)
	if err != nil {
		return 0, fmt.Errorf("login failed at Encrypt err %w", err)
	}

	user, err := a.userRepo.GetUserByEmail(encEmail)
	if err != nil {
		return 0, fmt.Errorf("login failed at GetUserByEmail err %w", err)
	}

	pass, err := a.userRepo.Decrypt(user.Password)
	if err != nil {
		return 0, fmt.Errorf("login failed at Decrypt err %w", err)
	}

	if pass != fv.Password {
		return 0, fmt.Errorf("login failed at compare pass err %w", err)
	}

	claims := &jwtCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(a.config.JwtSecret))
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

	return user.ID, nil
}

func (a *AuthUseCase) IdentifyJwtUser(id int) error {
	_, err := a.userRepo.GetUserById(id)
	if err != nil {
		return fmt.Errorf("IdentifyJwtUser err %w", err)
	}

	return nil
}

func (a *AuthUseCase) DeleteCookie(c echo.Context, ck *http.Cookie) {
	ck.Value = ""
	ck.MaxAge = -1
	c.SetCookie(ck)
}
