package entity

import (
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	AccountID  int `gorm:"column:account_id"`
	Name       string
	ID         int
	Kind       string
	LocationID int
	ContactID  int
}
