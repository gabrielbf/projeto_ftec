package entity

import (
	"gorm.io/gorm"
)

const (
	created string = "CREATED"

	user    string = "USER"
	partner string = "PARTNER"
)

type accountStatusEnum struct {
	Created string
}

var AccountStatusEnum = accountStatusEnum{
	Created: created,
}

type accountTypeEnum struct {
	User    string
	Partner string
}

var AccountTypeEnum = accountTypeEnum{
	User:    user,
	Partner: partner,
}

type Account struct {
	gorm.Model
	Email    string
	Password string
	Type     string
	Status   string
}
