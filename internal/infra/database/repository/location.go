package repository

import (
	"context"

	"github.com/ftec-project/internal/domain/location"
	"github.com/ftec-project/internal/infra/database/entity"
	"gorm.io/gorm"
)

type locationRepositoryImpl struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) location.Repository {
	return &locationRepositoryImpl{
		db: db,
	}
}

func (r locationRepositoryImpl) Create (ctx context.Context, location *entity.Location) error {
	return r.db.
		WithContext(ctx).
		Table("location").
		Create(&location).
		Error
}