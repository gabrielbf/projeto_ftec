package sample

import (
	"context"

	"github.com/sample-project/internal/infra/database/entity"
)

type Service interface {
	CreateSample(ctx context.Context, createDTO CreateDTO) (entity.Sample, error)
	GetByReferenceUUID(ctx context.Context, referenceUUID string) (entity.Sample, error)
}

type CreateDTO struct {
	ReferenceUUID string
}
