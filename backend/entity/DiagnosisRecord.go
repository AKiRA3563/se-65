package entity

import (
	"time"
	"gorm.io/gorm"
)

type Disease struct {
	
	gorm.Model

	Name			string
	Description		string

	DiagnosisRecords	[]DiagnosisRecord	`gorm:"foreignKey:Disease_ID"`
}

type DiagnosisRecord struct {

	gorm.Model

	PatientID			*uint
	Patient				PatientRegistered	`gorm:"references:ID"`

	DoctorID			*uint
	Doctor				Employee			`gorm:"references:ID"`

	HistorySheetID		*uint
	HistorySheet		HistorySheet		`gorm:"references:ID"`

	DiseaseID			*uint
	Disease				Disease				`gorm:"references:ID"`
	
	Examination			string
	MedicalCertificate 	bool	
	Date				time.Time

}