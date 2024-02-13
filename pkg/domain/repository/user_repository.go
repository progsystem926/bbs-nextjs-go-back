package repository

import "github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model/graph"

type User interface {
	GetUserById(id string) (*graph.User, error)
	GetUserByEmail(email string) (*graph.User, error)
	CreateUser(user *graph.User) (*graph.User, error)
	Encrypt(plain string) (string, error)
	Decrypt(encrypted string) (string, error)
}
