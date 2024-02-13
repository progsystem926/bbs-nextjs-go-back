package resolver

import "github.com/progsystem926/bbs-nextjs-go-back/pkg/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostUseCase usecase.Post
	UserUseCase usecase.User
}
