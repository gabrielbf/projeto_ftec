package location

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreateLocation(ctx context.Context, createDTO CreateDTO) (entity.Location, error)
}

type CreateDTO struct {
	Country string
	State string
	City string
	Neighborhood string
	Street string
	Number string
	Complement string
	Cep string
	Coordinates string
	Description string
}