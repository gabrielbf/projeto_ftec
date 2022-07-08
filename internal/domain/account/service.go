package account

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreateAccount(ctx context.Context, createDTO CreateDTO) (entity.Account, error)
	GetByReferenceUUID(ctx context.Context, referenceUUID string) (entity.Account, error)
}

type CreateDTO struct {
	Email    string
	Password string
}
