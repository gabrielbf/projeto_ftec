package response

import (
	"time"

	"github.com/ftec-project/internal/infra/database/entity"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	FirstName string
	LastName  string
	AccountID int
	CreatedAt time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the user ressource was saved
	UpdatedAt time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the user ressource was updated
	DeletedAt null.Time `json:"deleted_at" extensions:"x-nullable" example:"null" swaggertype:"string"` // Date time when the user ressource was deleted
}

func (c *User) FromUser(user entity.User) {
	c.FirstName = user.FirstName
	c.LastName = user.LastName
	c.AccountID = user.AccountID
	c.CreatedAt = user.CreatedAt
	c.UpdatedAt = user.UpdatedAt
	// c.DeletedAt = user.DeletedAt
}
