package pet

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Repository interface {
	Create(ctx context.Context, account *entity.Pet) error
}