package repository

import (
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
)

type User interface {
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	Encrypt(plain string) (string, error)
	Decrypt(encrypted string) (string, error)
}
