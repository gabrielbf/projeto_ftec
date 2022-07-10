package service

import (
	"context"
	"errors"

	"github.com/ftec-project/internal/domain/pet"
	"github.com/ftec-project/internal/domain/location"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type petServiceImpl struct {
	petRepository pet.Repository
	locatioService location.Service
}

func NewPetService(petRepository pet.Repository, locationService location.Service) pet.Service {
	return &petServiceImpl{
		petRepository: petRepository,
		locationService: locatioService,
	}
}

func (s petServiceImpl) CreatePet(ctx context.Context, createDTO pet.CreateDto) (entity.Pet, error) {
	// Cria location
	//location, err := s.locationService.CreateLocation(ctx, )
}