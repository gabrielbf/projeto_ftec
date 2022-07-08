package repository

import (
	"context"

	"github.com/ftec-project/internal/domain/sample"
	"github.com/ftec-project/internal/infra/database/entity"
	"gorm.io/gorm"
)

type sampleRepositoryImpl struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) sample.Repository {
	return &sampleRepositoryImpl{
		db: db,
	}
}

func (r sampleRepositoryImpl) Create(ctx context.Context, sample *entity.Sample) error {
	return r.db.
		WithContext(ctx).
		Table("samples").
		Create(&sample).
		Error
}

func (r sampleRepositoryImpl) Update(ctx context.Context, sample *entity.Sample) error {
	return r.db.WithContext(ctx).
		Table("samples").
		Updates(&sample).
		Error
}

func (r sampleRepositoryImpl) GetByReferenceUUID(ctx context.Context, sample *entity.Sample, reference string) error {
	return r.db.
		WithContext(ctx).
		Table("samples").
		Where("reference_uuid = ?", reference).
		First(sample).
		Error
}
