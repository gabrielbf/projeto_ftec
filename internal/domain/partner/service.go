package partner

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreatePartner(ctx context.Context, createDTO CreateDTO) (entity.Partner, error)
}

type CreateDTO struct {
	Name string
}