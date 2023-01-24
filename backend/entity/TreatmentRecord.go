package entity

import (
	"time"
	"gorm.io/gorm"
)

type TreatmentRecord struct {
	
	gorm.Model

	PatientID			*uint
	Patient				PatientRegistered	`gorm:"references:ID"`

	DoctorID			*uint
	Doctor				Employee			`gorm:"references:ID"`
	
	DiagnosisRecordID	*uint
	DiagnosisRecord		DiagnosisRecord		`gorm:"references:ID"`

	Treatment			string
	Note				string
	
	MedicineRecordID	*uint
	MedicineRecord		MedicineRecord		`gorm:"reference:ID"`
	MedicineQuantity	int

	Date				time.Time
	
}