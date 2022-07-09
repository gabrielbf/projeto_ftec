package response

import (
	"fmt"
	"time"

	"github.com/ftec-project/internal/infra/database/entity"
	"gopkg.in/guregu/null.v4"
)

type Account struct {
	ID        string    `json:"id" example:"1"`
	Email     string    `json:"email" example:"user@email.com"`
	Password  string    `json:"password" example:"password"`
	Status    string    `json:"status" example:"CREATED"`
	CreatedAt time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the account ressource was saved
	UpdatedAt time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the account ressource was updated
	DeletedAt null.Time `json:"deleted_at" extensions:"x-nullable" example:"null" swaggertype:"string"` // Date time when the account ressource was deleted
}

func (c *Account) FromAccount(account entity.Account) {
	c.ID = fmt.Sprint(account.Model.ID)
	c.Email = account.Email
	c.Password = account.Password
	c.Status = account.Status
	c.CreatedAt = account.CreatedAt
	c.UpdatedAt = account.UpdatedAt
	// c.DeletedAt = account.DeletedAt
}
