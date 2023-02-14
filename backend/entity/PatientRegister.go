package entity

import (
	"gorm.io/gorm"
)

type PatientRegister struct {
	gorm.Model

	//Prefix		string
	FirstName	string
	LastName	string

	GenderID	*uint
	Gender		Gender	`gorm:"references:ID"`
	
	Age			int
	Mobile		string	`gorm:"uniqueIndex"`
	Email		string	`gorm:"uniqueIndex"`

	HistorySheets	[]HistorySheet	`gorm:"foreignKey:PatientRegisterID"`
	// DiagnosisRecords	[]DiagnosisRecord	`gorm:"foreignKey:PatientRegisterID"`		
	// TreatmentRecords	[]TreatmentRecord	`gorm:"foreignKey:PatientRegisterID"`
}