package entity

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Medicine struct {
	gorm.Model

	Name        string
	Description string
	Price       float32

	MedicineOrder []MedicineOrder `gorm:"foreignKey:MedicineID"`
}

type MedicineOrder struct {
	gorm.Model

	MedicineID *uint
	Medicine   Medicine `gorm:"references:ID" valid:"-"`

	OrderAmount int `valid:"int, range(0|100)~Order Amount must not be negative"`

	TreatmentRecordID *uint
	TreatmentRecord   TreatmentRecord `gorm:"references:ID" valid:"-"`
}

type TreatmentRecord struct {
	gorm.Model

	//DoctorID เป็น FK
	DoctorID *uint
	Doctor   Employee `gorm:"references:ID" valid:"-"`

	//DiagnosisRecord เป็น FK
	DiagnosisRecordID *uint
	DiagnosisRecord   DiagnosisRecord `gorm:"references:ID" valid:"-"`

	Treatment   string `valid:"required~Treatment cannot be blank"`
	Note        string `valid:"-"`
	Appointment *bool

	MedicineOrders []MedicineOrder `gorm:"foreignKey:TreatmentRecordID;  constraint:OnDelete:CASCADE" valid:"-"`

	Date time.Time `valid:"present~Date must be present"`
}

func init() {

	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(time.Hour*-12)) && t.Before(time.Now().Add(time.Hour*12))
	})

}

func BooleanNotNull(t *bool) (bool, error) {
	if t == nil {
		return false, fmt.Errorf("Appointment cannot be Null")
	} else {
		return true, nil
	}
}
