package account

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Service interface {
	CreateAccount(ctx context.Context, createDTO CreateDTO) (entity.Account, entity.Contact, entity.User, entity.Partner, error)
}

type CreateDTO struct {
	Email        string
	Password     string
	Type         string
	Name         string
	PhoneNumber  string
	Country      string
	State        string
	City         string
	Neighborhood string
	Street       string
	Number       string
}
