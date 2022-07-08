package response

import (
	"time"

	"github.com/ftec-project/internal/infra/database/entity"
	"gopkg.in/guregu/null.v4"
)

type Partner struct {
	Name      string
	AccountID int
	CreatedAt time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the partner  ressource was saved
	UpdatedAt time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the partner  ressource was updated
	DeletedAt null.Time `json:"deleted_at" extensions:"x-nullable" example:"null" swaggertype:"string"` // Date time when the partner  ressource was deleted
}

func (c *Partner) FromPartner(partner entity.Partner) {
	c.AccountID = partner.AccountID
	c.Name = partner.Name
	c.CreatedAt = partner.CreatedAt
	c.UpdatedAt = partner.UpdatedAt
	// c.DeletedAt = partner .DeletedAt
}
