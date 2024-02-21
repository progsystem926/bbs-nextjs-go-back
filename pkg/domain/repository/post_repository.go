package repository

import "github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"

type Post interface {
	GetPosts() ([]*model.Post, error)
	CreatePost(post *model.Post) (*model.Post, error)
	DeletePost(post *model.Post) (bool, error)
}
