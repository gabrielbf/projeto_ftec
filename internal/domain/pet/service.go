package pet

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreatePet(ctx context.Context, createDTO CreateDTO) (entity.Pet, error)
}

type CreateDTO struct {
	Descritption string
	Species	string
	Status string
	//Location entity.Location
	Country string
	State string
	City string
	Neighborhood string
	Street string
	Number string
	Complement string
	Cep string
	Coordinates string
	LocationDescription string
}
