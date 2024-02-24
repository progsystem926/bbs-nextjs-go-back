package usecase

import (
	"fmt"
	"time"

	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/repository"
)

type Post interface {
	GetPosts() ([]*model.Post, error)
	CreatePost(text *string, userId int) (*model.Post, error)
	DeletePost(id int) (bool, error)
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

func (p *PostUseCase) CreatePost(text *string, userId int) (*model.Post, error) {
	timestamp := time.Now().Format(time.DateTime)

	post := model.Post{
		Text:      *text,
		UserID:    userId,
		CreatedAt: timestamp,
	}

	created, err := p.postRepo.CreatePost(&post)
	if err != nil {
		return nil, fmt.Errorf("useCase CreatePost() err: %w", err)
	}

	return created, nil
}

func (p *PostUseCase) DeletePost(id int) (bool, error) {
	deleted, err := p.postRepo.DeletePost(&model.Post{ID: id})
	if err != nil {
		return false, fmt.Errorf("useCase DeletePost() err: %w", err)
	}

	return deleted, nil
}
