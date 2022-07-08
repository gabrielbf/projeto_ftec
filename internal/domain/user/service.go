package user

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreateUser(ctx context.Context, createDTO CreateDTO) (entity.User, error)
}

type CreateDTO struct {
	FirstName string
	LastName  string
	AccountID int
}
