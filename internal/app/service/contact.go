package service

import (
	"context"

	"github.com/ftec-project/internal/domain/contact"
	"github.com/ftec-project/internal/infra/database/entity"
	"github.com/rs/zerolog/log"
)

type contactServiceImpl struct {
	contactRepository contact.Repository
}

func NewContactService(contactRepository contact.Repository) contact.Service {
	return &contactServiceImpl{
		contactRepository: contactRepository,
	}
}

func (s contactServiceImpl) CreateContact(ctx context.Context, createDTO contact.CreateDTO) (entity.Contact, error) {
	log.Info().Msg("creating contacts")

	var emailContact entity.Email
	var phoneContact entity.Phone

	contact, err := s.contactRepository.Create(ctx)
	if err != nil {
		log.Error().Interface("Contact", contact).Msg("Error when saving Contact")
		return entity.Contact{}, err
	}

	if createDTO.Email.Email != "" {
		log.Info().Msg("creating email")
		emailContact = entity.Email{
			Email: createDTO.Email.Email,
		}
		err := s.contactRepository.CreateEmail(ctx, contact, &emailContact)
		if err != nil {
			log.Error().Interface("Email", emailContact).Msg("Error when saving email")
			return entity.Contact{}, err
		}
	}

	if createDTO.PhoneNumber.Number != "" {
		log.Info().Msg("creating phone number")
		phoneContact = entity.Phone{
			PhoneNumber: createDTO.PhoneNumber.Number,
		}
		err := s.contactRepository.CreatePhoneNumber(ctx, contact, &phoneContact)
		if err != nil {
			log.Error().Interface("Phone", phoneContact).Msg("Error when saving phone")
			return entity.Contact{}, err
		}
	}

	return entity.Contact{
		ID:    contact.ID,
		Email: emailContact,
		Phone: phoneContact,
	}, nil
}
