package infra

import (
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
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

func (r *userRepository) GetUserById(id int) (*model.User, error) {
	var user *model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, xerrors.Errorf("get users by id failed , %w", err)
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user *model.User
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

func (r *userRepository) CreateUser(user *model.User) (*model.User, error) {
	if result := r.db.Create(user); result.Error != nil {
		return nil, xerrors.Errorf("repository CreateUser() err %w", result.Error)
	}

	return user, nil
}
