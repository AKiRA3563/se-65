package entity

import (
	"gorm.io/gorm"
)

type HistorySheet struct {
	gorm.Model

	PatientRegisteredID	*uint
	PatientRegistered	PatientRegistered	`gorm:"references:ID"`

	NurseID		*uint
	Nurse		Employee	`gorm:"references:ID"`

	Weight			float32
	Height			float32
	Temperature		float32
	
}