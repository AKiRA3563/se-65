package entity

import (

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

)

var db *gorm.DB

func DB() *gorm.DB {
	
	return db

}

func SetuoDatabase () {

	database, err := gorm.Open(sqlite.Open("se-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&DiagnosisRecord{},
		&TreatmentRecord{},
		&HistorySheet{},
		&MedicineRecord{},
		&PatientRegistered{},
		&TreatmentRecord{},
	)

	db = database

}
