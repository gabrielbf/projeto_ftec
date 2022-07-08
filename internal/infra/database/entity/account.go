package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ReferenceUUID string
}
