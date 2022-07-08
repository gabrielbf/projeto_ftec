package user

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreateUser(ctx context.Context, createDTO CreateDTO) (entity.User, error)
	GetByReferenceUUID(ctx context.Context, referenceUUID string) (entity.User, error)
}

type CreateDTO struct {
	ReferenceUUID string
}
