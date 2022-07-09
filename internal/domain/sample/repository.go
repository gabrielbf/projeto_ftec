package sample

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Repository interface {
	Create(ctx context.Context, sample *entity.Sample) error
	Update(ctx context.Context, sample *entity.Sample) error
	GetByReferenceUUID(ctx context.Context, sample *entity.Sample, reference string) error
}
