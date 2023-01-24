package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	//Title		string
	Name		string
	gender		string
	PhoneNumber	string	`gorm:"uniqueIndex"`
	Email		string	`gorm:"uniqueIndex"`
	Password	string

}