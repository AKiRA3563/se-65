package entity

import (
	// "fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestTreatmentRecordPass(t *testing.T) {
	g := NewGomegaWithT(t)

	treatment := TreatmentRecord{
		Note:             "yes",
		Treatment:        "yes",
		MedicineQuantity: 20,
		Date:             time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(treatment)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())

}

func TestTreatmentNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	treatment := TreatmentRecord{
		Note:             "",
		Treatment:        "", //wrong
		MedicineQuantity: 20,
		Date:             time.Now(),
	}

	// ตรวจสอบด้วย govalidation
	ok, err := govalidator.ValidateStruct(treatment)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าจ้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Treatment cannot be blank"))
}

func TestMedQuantNotNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	treatment := TreatmentRecord{
		Note:             "",
		Treatment:        "yes",
		MedicineQuantity: -56,
		Date:             time.Now(),
	}

	// ตรวจสอบด้วย govalidation
	ok, err := govalidator.ValidateStruct(treatment)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าจ้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("MedicineQuantity must not be negative"))

}

func TestDateNotbePastandFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []time.Time{
		time.Now().Add(24 * time.Hour),
		time.Now().Add(-24 * time.Hour),
	}

	for _, fixture := range fixtures {
		treatment := TreatmentRecord{
			Note:             "",
			Treatment:        "yes",
			MedicineQuantity: 9,
			Date:             fixture,
		}

		// ตรวจสอบด้วย govalidation
		ok, err := govalidator.ValidateStruct(treatment)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าจ้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Date must be present"))

	}
}
