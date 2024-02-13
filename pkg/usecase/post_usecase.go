package usecase

import (
	"fmt"
	"time"

	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/repository"
)

type Post interface {
	GetPosts() ([]*model.Post, error)
	CreatePost(text *string, userId *string) (*model.Post, error)
}

type PostUseCase struct {
	postRepo repository.Post
}

func NewPostUseCase(postRepo repository.Post) Post {
	PostUseCase := PostUseCase{postRepo: postRepo}
	return &PostUseCase
}

func (p *PostUseCase) GetPosts() ([]*model.Post, error) {
	posts, err := p.postRepo.GetPosts()
	if err != nil {
		return nil, fmt.Errorf("useCase GetPosts() err: %w", err)
	}

	return posts, nil
}

func (p *PostUseCase) CreatePost(text *string, userId *string) (*model.Post, error) {
	timestamp := time.Now().Format("2024-01-01 01:01:01")

	post := model.Post{
		Text:      *text,
		UserID:    *userId,
		CreatedAt: timestamp,
	}

	created, err := p.postRepo.CreatePost(&post)
	if err != nil {
		return nil, fmt.Errorf("useCase CreatePost() err: %w", err)
	}

	return created, nil
}
