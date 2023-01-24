package entity

import (
	"gorm.io/gorm"
)

type MedicineRecord struct {
	
	gorm.Model

	Name		string
	Description	string
	Quantity 	int
}