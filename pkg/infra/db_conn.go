package infra

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnector(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, xerrors.Errorf("db connection failed:%w", err)
	}

	return db, nil
}