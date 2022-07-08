package partner

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Repository interface {
	Create(ctx context.Context, partner *entity.Partner) error
	Update(ctx context.Context, partner *entity.Partner) error
	GetByReferenceUUID(ctx context.Context, partner *entity.Partner, reference string) error
}
