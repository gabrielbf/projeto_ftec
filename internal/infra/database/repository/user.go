package repository

import (
	"context"

	"github.com/ftec-project/internal/domain/user"
	"github.com/ftec-project/internal/infra/database/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r userRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	return r.db.
		WithContext(ctx).
		Table("_user").
		Create(&user).
		Error
}
