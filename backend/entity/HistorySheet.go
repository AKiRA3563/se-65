package entity

import (
	"gorm.io/gorm"
)

type HistorySheet struct {
	gorm.Model
	Weight                 float32
	Height                 float32
	Temperature            float32
	SystolicBloodPressure  uint8
	DiastolicBloodPressure uint8
	HeartRate              int
	RespiratoryRate        uint8
	OxygenSaturation       uint8
	DrugAllergySymtom      string
	PatientSymtom          string

	PatientRegisterID *uint
	PatientRegister   PatientRegister `gorm:"references:ID"`

	NurseID *uint
	Nurse   Employee `gorm:"references:ID"`

	DrugAllergyID    *uint
	DrugAllergy      DrugAllergy       `gorm:"references:ID"`
	DiagnosisRecords []DiagnosisRecord `gorm:"foreignKey:HistorySheetID"`
}

type DrugAllergy struct {
	gorm.Model
	Name          string
	HistorySheets []HistorySheet `gorm:"foreignKey:DrugAllergyID"`
}
