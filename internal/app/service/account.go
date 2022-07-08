package service

import (
	"context"

	"github.com/ftec-project/internal/domain/account"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type accountServiceImpl struct {
	accountRepository account.Repository
}

func NewAccountService(accountRepository account.Repository) account.Service {
	return &accountServiceImpl{
		accountRepository: accountRepository,
	}
}

func (s accountServiceImpl) CreateAccount(ctx context.Context, createDTO account.CreateDTO) (entity.Account, error) {
	log.Info().Msg("saving account request")
	account := entity.Account{}

	// err := s.accountRepository.Create(ctx, &account)
	// if err != nil {
	// 	log.Error().Interface("Account", account).Msg("Error when saving account")
	// 	return entity.Account{}, err
	// }

	return account, nil
}
