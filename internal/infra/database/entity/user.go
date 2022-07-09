package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	AccountID int `gorm:"column:account_id"`
	ID        int
	ContactId int
}
