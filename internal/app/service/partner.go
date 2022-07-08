package service

import (
	"context"

	"github.com/ftec-project/internal/domain/partner"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type partnerServiceImpl struct {
	partnerRepository partner.Repository
}

func NewPartnerService(partnerRepository partner.Repository) partner.Service {
	return &partnerServiceImpl{
		partnerRepository: partnerRepository,
	}
}

func (s partnerServiceImpl) CreatePartner(ctx context.Context, createDTO partner.CreateDTO) (entity.Partner, error) {
	log.Info().Msg("saving partner request")
	partner := entity.Partner{}

	// err := s.partnerRepository.Create(ctx, &partner)
	// if err != nil {
	// 	log.Error().Interface("Partner", partner).Msg("Error when saving partner")
	// 	return entity.Partner{}, err
	// }

	return partner, nil
}
