package contact

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreateContact(ctx context.Context, createDTO CreateDTO) (entity.Contact, error)
}

type CreateDTO struct {
	PhoneNumber PhoneNumber
	Email       Email
}

type PhoneNumber struct {
	Number string
}

type Email struct {
	Email string
}
