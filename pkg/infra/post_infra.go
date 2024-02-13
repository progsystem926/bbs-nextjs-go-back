package infra

import (
	"golang.org/x/xerrors"

	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/repository"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.Post {
	return &postRepository{db}
}

func (r *postRepository) GetPosts() ([]*model.Post, error) {
	var records []model.Post
	if result := r.db.Find(&records); result.Error != nil {
		return nil, xerrors.Errorf("repository GetMessages() err %w", result.Error)
	}

	var res []*model.Post
	for _, record := range records {
		record := record
		res = append(res, &record)
	}

	return res, nil
}

func (r *postRepository) CreatePost(post *model.Post) (*model.Post, error) {
	if result := r.db.Create(post); result.Error != nil {
		return nil, xerrors.Errorf("repository CreatePost() err %w", result.Error)
	}

	return post, nil
}
