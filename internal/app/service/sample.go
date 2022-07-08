package service

import (
	"context"

	"github.com/ftec-project/internal/domain/sample"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type sampleServiceImpl struct {
	sampleRepository sample.Repository
}

func NewSampleService(sampleRepository sample.Repository) sample.Service {
	return &sampleServiceImpl{
		sampleRepository: sampleRepository,
	}
}

func (s sampleServiceImpl) CreateSample(ctx context.Context, createDTO sample.CreateDTO) (entity.Sample, error) {
	log.Info().Msg("saving sample request")
	sample := entity.Sample{
		ReferenceUUID: createDTO.ReferenceUUID,
	}

	err := s.sampleRepository.Create(ctx, &sample)
	if err != nil {
		log.Error().Interface("Sample", sample).Msg("Error when saving sample")
		return entity.Sample{}, err
	}

	return sample, nil
}

func (s sampleServiceImpl) GetByReferenceUUID(ctx context.Context, referenceUUID string) (sample entity.Sample, err error) {
	return sample, s.sampleRepository.GetByReferenceUUID(ctx, &sample, referenceUUID)
}
