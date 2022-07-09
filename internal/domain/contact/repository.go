package contact

import (
	"context"

	"github.com/ftec-project/internal/infra/database/entity"
)

type Repository interface {
	Create(ctx context.Context) (*entity.Contact, error)
	CreateEmail(ctx context.Context, contact *entity.Contact, email *entity.Email) error
	CreatePhoneNumber(ctx context.Context, contact *entity.Contact, phoneNumber *entity.Phone) error
}
