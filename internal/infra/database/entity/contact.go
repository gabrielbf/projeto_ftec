package entity

import "gorm.io/gorm"

type Contact struct {
	ID    int   `gorm:"primarykey"`
	Email Email `gorm:"-"`
	Phone Phone `gorm:"-"`
}

type Email struct {
	gorm.Model
	ContactID   int
	Email       string
	Description string
	Status      string
}

type Phone struct {
	gorm.Model
	ContactID   int
	PhoneNumber string
	Description string
	Status      string
}
