package service

import (
	"context"
	"errors"

	"github.com/ftec-project/internal/domain/account"
	"github.com/ftec-project/internal/domain/contact"
	"github.com/ftec-project/internal/domain/partner"
	"github.com/ftec-project/internal/domain/user"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type accountServiceImpl struct {
	accountRepository account.Repository
	contactService    contact.Service
	userService       user.Service
	partnerService    partner.Service
}

func NewAccountService(accountRepository account.Repository, contactService contact.Service, userService user.Service, partnerService partner.Service) account.Service {
	return &accountServiceImpl{
		accountRepository: accountRepository,
		contactService:    contactService,
		userService:       userService,
		partnerService:    partnerService,
	}
}

func (s accountServiceImpl) CreateAccount(ctx context.Context, createDTO account.CreateDTO) (entity.Account, entity.Contact, entity.User, entity.Partner, error) {
	var createdUser entity.User
	var createdPartner entity.Partner

	// 1 CRIA CONTATO
	contact, err := s.contactService.CreateContact(ctx, contact.CreateDTO{
		PhoneNumber: contact.PhoneNumber{
			Number: createDTO.PhoneNumber,
		},
		Email: contact.Email{
			Email: createDTO.Email,
		},
	})
	if err != nil {
		return entity.Account{}, entity.Contact{}, entity.User{}, entity.Partner{}, err
	}

	log.Info().Msg("saving account request")
	account := entity.Account{
		Email:    createDTO.Email,
		Password: createDTO.Password,
		Type:     createDTO.Type,
		Status:   entity.AccountStatusEnum.Created,
	}

	// 2 CRIA CONTA
	err = s.accountRepository.Create(ctx, &account)
	if err != nil {
		log.Error().Interface("Account", account).Msgf("Error when saving account %s", err)
		return entity.Account{}, entity.Contact{}, entity.User{}, entity.Partner{}, err
	}

	// 3 CRIA LOCALIZACAO

	// 4 CRIA PARCEIRO
	if createDTO.Type == entity.AccountTypeEnum.Partner {
		createdPartner, err = s.partnerService.CreatePartner(ctx, partner.CreateDTO{
			AccountID: int(account.ID),
			Name:      createDTO.Name,
			// LocationID: location.ID,
			ContactID: contact.ID,
		})
		if err != nil {
			return entity.Account{}, entity.Contact{}, entity.User{}, entity.Partner{}, err
		}

		// 4 CRIA USUARIO
	} else if createDTO.Type == entity.AccountTypeEnum.User {
		createdUser, err = s.userService.CreateUser(ctx, user.CreateDTO{
			AccountID: int(account.ID),
			FirstName: createDTO.Name,
			ContactID: contact.ID,
		})
		if err != nil {
			return entity.Account{}, entity.Contact{}, entity.User{}, entity.Partner{}, err
		}
	} else {
		INVALID_ACCOUNT_TYPE := errors.New("invalid account type")
		log.Err(INVALID_ACCOUNT_TYPE)
		return entity.Account{}, entity.Contact{}, entity.User{}, entity.Partner{}, INVALID_ACCOUNT_TYPE
	}

	return account, contact, createdUser, createdPartner, nil
}
