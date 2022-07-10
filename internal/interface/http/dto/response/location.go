package response

import (
	"time"

	"github.com/ftec-project/internal/infra/database/entity"
	"gopkg.in/guregu/null.v4"
)

type Location struct {
	Country      string
	State        string
	City         string
	Neighborhood string
	Street       string
	Number       string
	Complement   string
	Cep          string
	Coordinates  string
	Description  string
	CreatedAt time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the user ressource was saved
	UpdatedAt time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z"`                              // Date time when the user ressource was updated
	DeletedAt null.Time `json:"deleted_at" extensions:"x-nullable" example:"null" swaggertype:"string"` // Date time when the user ressource was deleted
}

func (c *Location) FromLocation(location entity.Location) {
	c.Country = location.Country
	c.State = location.State
	c.City = location.City
	c.Neighborhood = location.Neighborhood
	c.Street = location.Street
	c.Number = location.Number
	c.Complement = location.Complement
	c.Cep = location.Cep
	c.Coordinates = location.Coordinates
	c.Description = location.Description
}