package entity

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Disease struct {
	gorm.Model

	Name        string
	Description string

	DiagnosisRecords []DiagnosisRecord `gorm:"foreignKey:DiseaseID"`
}

type DiagnosisRecord struct {
	gorm.Model

	//DoctorID เป็น FK
	DoctorID *uint
	Doctor   Employee `gorm:"references:ID" valid:"-"`

	//HistorySheetID เป็น FK
	HistorySheetID *uint
	HistorySheet   HistorySheet `gorm:"references:ID" valid:"-"`

	//DiseaseID เป็น FK
	DiseaseID *uint
	Disease   Disease `gorm:"references:ID" valid:"-"`

	Examination        string `valid:"required~Examination cannot be Blank"`
	MedicalCertificate *bool
	Date               time.Time `valid:"present~Date must be present"`

	TreatmentRecords []TreatmentRecord `gorm:"foreignKey:DiagnosisRecordID; constraint:OnUpdate:CASCADE"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(time.Hour*-12)) && t.Before(time.Now().Add(time.Hour*12))
	})

}

// ใช้แค่ validate backend เท่านั้น
func CheckBool(t *bool) (bool, error) {
	if t == nil {
		return false, fmt.Errorf("MedicalCertificate cannot be Null")
	} else {
		return true, nil
	}
}
