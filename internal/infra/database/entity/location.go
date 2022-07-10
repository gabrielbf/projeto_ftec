package entity

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	//id           int
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
}
