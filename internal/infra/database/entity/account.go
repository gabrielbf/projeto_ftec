package entity

import (
	"gorm.io/gorm"
)

const (
	created string = "CREATED"
)

type accountStatusEnum struct {
	Created string
}

var AccountStatusEnum = accountStatusEnum{
	Created: created,
}

type Account struct {
	gorm.Model
	Email    string
	Password string
	Status   string
}
