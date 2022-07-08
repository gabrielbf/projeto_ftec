package entity

import (
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	ReferenceUUID string
}
