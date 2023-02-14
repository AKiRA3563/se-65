package entity

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	// "github.com/go-playground/validator/v10"
)

type Disease struct {
	gorm.Model

	Name        string
	Description string

	DiagnosisRecords []DiagnosisRecord	`gorm:"foreignKey:DiseaseID"`
}

type MedicalCertificate	struct {
	gorm.Model

	Label string

	DiagnosisRecords	[]DiagnosisRecord	`gorm:"foreignKey:MedicalCertificateID"`
}

type DiagnosisRecord struct {
	gorm.Model

	// PatientRegisterID *uint
	// PatientRegister   PatientRegister `gorm:"references:id" valid:"-"`

	//DoctorID เป็น FK
	DoctorID *uint
	Doctor   Employee `gorm:"references:id" valid:"-"`

	//HistorySheetID เป็น FK
	HistorySheetID *uint
	HistorySheet   HistorySheet `gorm:"references:id" valid:"-"`

	//DiseaseID เป็น FK
	DiseaseID *uint
	Disease   Disease `gorm:"references:id" valid:"-"`

	//MedCerID เป็น FK
	// MedicalCertificateID	*uint				//`valid:"int,required~MedicalCertificate must not Null"`
	// MedicalCertificate		MedicalCertificate	`gorm:"references:id"`

	Examination        string	`valid:"required~Examination cannot be Blank"`
	MedicalCertificate	*bool
	Date               time.Time `valid:"present~Recording time must be current"`

	// TreatmentRecord []TreatmentRecord `gorm:"foreignKey:DiagnosisRecordID"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		// ปัจจุบัน บวกลบไม่เกิน 3 นาที		
		if t.Before(time.Now().Add(-3*time.Minute)) || t.After(time.Now().Add(3*time.Minute)) {
			return false
		} else {
			return true
		}
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
