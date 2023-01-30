package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type DiagnosisRecord struct {
	gorm.Model
	PatientID	int
	DoctorID	int
	HistorySheetID	int
	DiseaseID	int

	Examination	string	`valid:"required~Examination cannot be blank"`
	MedicalCertificate	*bool	`valid:"null"`
	Date	time.Time
}

// ตรวจสอบค่าว่างของการตรวจโรคแล้วต้องเจอ Error
func TestExaminationNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	tt := true

	diagnosis := DiagnosisRecord{
		PatientID: 1,
		DoctorID: 1,
		HistorySheetID: 1,
		DiseaseID: 103,

		Examination: "", //wrong
		MedicalCertificate: &tt,
		Date: time.Now(),
	}

	// ตรวจสอบด้วย govalidation
	ok, err := govalidator.ValidateStruct(diagnosis)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())
	
	// err ต้องไม่เป็นค่า nil แปลว่าจ้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Examination cannot be blank"))
	
}

func TestMedicalCertificationMustBeChecked(t *testing.T) {
	g := NewGomegaWithT(t)

	tt := false

	diagnosis := DiagnosisRecord{
		PatientID: 1,
		DoctorID: 1,
		HistorySheetID: 1,
		DiseaseID: 103,

		Examination: "Headache",
		MedicalCertificate: &tt,
		Date: time.Now(),
	}

	// ตรวจสอบด้วย govalidation
	ok, err := govalidator.ValidateStruct(diagnosis)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())
	
	// err ต้องไม่เป็นค่า nil แปลว่าจ้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("MedicalCertification must be checked"))
}