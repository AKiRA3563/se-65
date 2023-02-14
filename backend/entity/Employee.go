package entity

import (
	"time"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FirstName string 
	LastName  string
	Email     string
	Password  string
	Employees []Employee `gorm:"foreignKey:AdminID"`
}

type Title struct {
	gorm.Model
	Name      string
	Employees []Employee `gorm:"foreignKey:TitleID"`
}

type Role struct {
	gorm.Model
	Name string
	Employees    []Employee `gorm:"foreignKey:RoleID"`
}

type Gender struct {
	gorm.Model
	Name      string
	Employees []Employee `gorm:"foreignKey:GenderID"`
	PatientRegisters	[]PatientRegister `gorm:"foreignKey:GenderID"`
}

type Employee struct {
	gorm.Model
	// //AdminID ทำหน้าที่ FK
	// AdminID *uint
	// Admin   Admin
	IDCard  string `gorm:"uniqueIndex"`
	// //TitleID ทำหน้าที่ FK
	// TitleID *uint
	// Title   Title
	FirstName    string	
	LastName	string
	//RoleID ทำหน้าที่ FK
	RoleID  *uint
	Role    Role `gorm:"references:ID"`

	Phonenumber string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Password    string 
	
	//GenderID ทำหน้าที่ FK
	GenderID *uint
	Gender   Gender `gorm:"references:ID"`
	//Salary 	uint32
	Birthday    time.Time

	DiagnosisRecords	[]DiagnosisRecord 		`gorm:"foreignKey:DoctorID"`
	HistorySheets		[]HistorySheet			`gorm:"foreignKey:NurseID"`
	TreatmentRecords	[]TreatmentRecord	`gorm:"foreignKey:DoctorID"`
}

