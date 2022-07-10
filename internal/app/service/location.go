package service

import (
	"context"

	"github.com/ftec-project/internal/domain/location"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type locationServiceImpl struct {
	locationRepository location.Repository
}

func NewLocationService(locationRepository location.Repository) location.Service {
	return &locationServiceImpl{
		locationRepository: locationRepository,
	}
}

func (s locationServiceImpl) CreateLocation(ctx context.Context, createDTO location.CreateDTO) (entity.Location, error) {
	log.Info().Msg("creating location")

	location := entity.Location{
		Country: createDTO.Country,
		State: createDTO.State,
		City: createDTO.City,
		Neighborhood: createDTO.Neighborhood,
		Street: createDTO.Street,
		Number: createDTO.Number,
		Complement: createDTO.Complement,
		Cep: createDTO.Cep,
		Coordinates: createDTO.Coordinates,
		Description: createDTO.Description,
	}

	err := s.locationRepository.Create(ctx, &location)
	if err != nil {
		log.Error().Interface("Location", location).Msgf("Error when saving account %s", err)
		return entity.Location{}, err
	}

	return location, nil
}