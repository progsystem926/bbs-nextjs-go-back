package usecase

import (
	"fmt"

	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/repository"
)

type User interface {
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(name *string, email *string, password *string) (*model.User, error)
}

type UserUseCase struct {
	userRepo repository.User
}

func NewUserUseCase(userRepo repository.User) User {
	UserUseCase := UserUseCase{userRepo: userRepo}
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

func (u *UserUseCase) CreateUser(name *string, email *string, password *string) (*model.User, error) {
	user := model.User{
		Name:     *name,
		Email:    *email,
		Password: *password,
	}

	created, err := u.userRepo.CreateUser(&user)
	if err != nil {
		return nil, fmt.Errorf("useCase CreateUser() err: %w", err)
	}

	return created, nil
}
