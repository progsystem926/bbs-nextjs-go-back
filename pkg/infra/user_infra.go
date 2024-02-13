package infra

import (
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model/graph"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/repository"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	config *config.Config
}

func NewUserRepository(db *gorm.DB, config *config.Config) repository.User {
	return &userRepository{db, config}
}

func (r *userRepository) GetUserById(id string) (*graph.User, error) {
	var user *graph.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, xerrors.Errorf("get users by id failed , %w", err)
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*graph.User, error) {
	var user *graph.User
	if err := r.db.Where("email =?", email).First(&user).Error; err != nil {
		return nil, xerrors.Errorf("get users by email failed, %w", err)
	}

	return user, nil
}

func (r *userRepository) Encrypt(plain string) (string, error) {
	var enc string
	err := r.db.Raw("SELECT HEX(AES_ENCRYPT(?, ?))", plain, r.config.EncryptKey).Scan(&enc).Error
	if err != nil {
		return enc, xerrors.Errorf("encrypt email failed , %w", err)
	}

	return enc, nil
}

func (r *userRepository) Decrypt(encrypted string) (string, error) {
	var dec string
	err := r.db.Raw("SELECT CONVERT(AES_DECRYPT(UNHEX(?), ?) USING utf8)", encrypted, r.config.EncryptKey).Scan(&dec).Error
	if err != nil {
		return dec, xerrors.Errorf("decrypt email failed, %w", err)
	}

	return dec, nil
}
