package response

import (
	"fmt"
	"time"

	"github.com/ftec-project/internal/infra/database/entity"
	"gopkg.in/guregu/null.v4"
)

type Account struct {
	ID          string    `json:"id" example:"1"`                                                         //
	Type        string    `json:"type" example:"USER"`                                                    //
	Status      string    `json:"status" example:"CREATED"`                                               //
	Email       string    `json:"email" example:"user@email.com"`                                         //
	PhoneNumber string    `json:"phone_number" example:"5551988888888"`                                   //
	Name        string    `json:"namer" example:"Lucas"`                                                  //
	CreatedAt   time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the account ressource was saved
	UpdatedAt   time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the account ressource was updated
	DeletedAt   null.Time `json:"deleted_at" extensions:"x-nullable" example:"null" swaggertype:"string"` // Date time when the account ressource was deleted
}

func (c *Account) FromAccount(account entity.Account, contact entity.Contact, user entity.User, partner entity.Partner, _type string) {
	c.ID = fmt.Sprint(account.Model.ID)
	c.Type = account.Type
	c.Status = account.Status
	c.Email = account.Email
	c.PhoneNumber = contact.Phone.PhoneNumber
	c.CreatedAt = account.CreatedAt
	c.UpdatedAt = account.UpdatedAt
	// c.DeletedAt = account.DeletedAt

	if _type == entity.AccountTypeEnum.Partner {
		c.Name = partner.Name
	} else {
		c.Name = user.FirstName
	}
}
