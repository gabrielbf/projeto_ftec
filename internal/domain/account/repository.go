package account

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Repository interface {
	Create(ctx context.Context, account *entity.Account) error
	Update(ctx context.Context, account *entity.Account) error
	GetByReferenceUUID(ctx context.Context, account *entity.Account, reference string) error
}
