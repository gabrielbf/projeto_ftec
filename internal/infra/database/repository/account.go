package repository

import (
	"context"

	"github.com/ftec-project/internal/domain/account"
	"github.com/ftec-project/internal/infra/database/entity"
	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) account.Repository {
	return &accountRepositoryImpl{
		db: db,
	}
}

func (r accountRepositoryImpl) Create(ctx context.Context, account *entity.Account) error {
	return r.db.
		WithContext(ctx).
		Table("accounts").
		Create(&account).
		Error
}

func (r accountRepositoryImpl) Update(ctx context.Context, account *entity.Account) error {
	return r.db.WithContext(ctx).
		Table("accounts").
		Updates(&account).
		Error
}

func (r accountRepositoryImpl) GetByReferenceUUID(ctx context.Context, account *entity.Account, reference string) error {
	return r.db.
		WithContext(ctx).
		Table("accounts").
		Where("reference_uuid = ?", reference).
		First(account).
		Error
}
