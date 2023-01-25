package entity

import (
	"gorm.io/gorm"
)

type PatientRegister struct {
	gorm.Model

	Prefix		string
	FirstName	string
	LastName	string
	Age			int
	Mobile		string	`gorm:"uniqueIndex"`
	Email		string	`gorm:"uniqueIndex"`

}