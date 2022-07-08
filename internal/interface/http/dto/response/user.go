package response

import (
	"time"

	"github.com/ftec-project/internal/infra/database/entity"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ReferenceUUID string    `json:"reference_uuid" example:"2e9948d6-5b0c-4984-a963-21e3ff0b2e25"`          // UUID sent by client and used as a client generated id. It will be the external reference for the requested user.
	CreatedAt     time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the user ressource was saved
	UpdatedAt     time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the user ressource was updated
	DeletedAt     null.Time `json:"deleted_at" extensions:"x-nullable" example:"null" swaggertype:"string"` // Date time when the user ressource was deleted
}

func (c *User) FromUser(user entity.User) {
	c.ReferenceUUID = user.ReferenceUUID
	c.CreatedAt = user.CreatedAt
	c.UpdatedAt = user.UpdatedAt
	// c.DeletedAt = user.DeletedAt
}
