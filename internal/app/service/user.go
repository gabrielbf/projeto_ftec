package service

import (
	"context"

	"github.com/ftec-project/internal/domain/user"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type userServiceImpl struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) user.Service {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (s userServiceImpl) CreateUser(ctx context.Context, createDTO user.CreateDTO) (entity.User, error) {
	log.Info().Msg("saving user request")
	user := entity.User{
		ReferenceUUID: createDTO.ReferenceUUID,
	}

	err := s.userRepository.Create(ctx, &user)
	if err != nil {
		log.Error().Interface("User", user).Msg("Error when saving user")
		return entity.User{}, err
	}

	return user, nil
}

func (s userServiceImpl) GetByReferenceUUID(ctx context.Context, referenceUUID string) (user entity.User, err error) {
	return user, s.userRepository.GetByReferenceUUID(ctx, &user, referenceUUID)
}
