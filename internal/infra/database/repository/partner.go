package repository

import (
	"context"

	"github.com/ftec-project/internal/domain/partner"
	"github.com/ftec-project/internal/infra/database/entity"
	"gorm.io/gorm"
)

type partnerRepositoryImpl struct {
	db *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) partner.Repository {
	return &partnerRepositoryImpl{
		db: db,
	}
}

func (r partnerRepositoryImpl) Create(ctx context.Context, partner *entity.Partner) error {
	return r.db.
		WithContext(ctx).
		Table("partner").
		Create(&partner).
		Error
}
