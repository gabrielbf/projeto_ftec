package repository

import (
	"context"

	"github.com/ftec-project/internal/domain/contact"
	"github.com/ftec-project/internal/infra/database/entity"
	"gorm.io/gorm"
)

type contactRepositoryImpl struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) contact.Repository {
	return &contactRepositoryImpl{
		db: db,
	}
}

func (r contactRepositoryImpl) Create(ctx context.Context) (*entity.Contact, error) {
	var contact *entity.Contact

	err := r.db.
		WithContext(ctx).
		Table("contact").
		Create(&contact).
		Error

	return contact, err
}

func (r contactRepositoryImpl) CreateEmail(ctx context.Context, contact *entity.Contact, email *entity.Email) error {
	email.ContactID = contact.ID
	return r.db.
		WithContext(ctx).
		Table("email").
		Create(&email).
		Error
}

func (r contactRepositoryImpl) CreatePhoneNumber(ctx context.Context, contact *entity.Contact, phone *entity.Phone) error {
	phone.ContactID = contact.ID
	return r.db.
		WithContext(ctx).
		Table("phone").
		Create(&phone).
		Error
}
