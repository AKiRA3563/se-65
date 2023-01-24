package controller

import (
	"net/http"

	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
)

// Post /TreamentRecord
func CreateTreatmentRecord(c *gin.Context) {

	var Patient			entity.PatientRegistered
	var Employee		entity.Employee
	var DiagnosisRecord entity.DiagnosisRecord
	var MedicineRecord	entity.MedicineRecord
	var TreatmentRecord	entity.TreatmentRecord

	
	if err := c.ShouldBindJSON(&TreatmentRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา Doctor ด้วย id
	if tx := entity.DB().Where("id = ?", TreatmentRecord.DoctorID).First(&Employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	//ค้นหา Patient ด้วย id
	if tx := entity.DB().Where("id = ?", TreatmentRecord.PatientID).First(&Patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	//ค้นหา MedicineRecord ด้วย id
	if tx := entity.DB().Where("id = ?", TreatmentRecord.MedicineRecordID).First(&MedicineRecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	//สร้าง TreamentRecord
	lr := entity.TreatmentRecord {
		Patient:			Patient,
		Doctor:				Employee,
		DiagnosisRecord:	DiagnosisRecord,
		MedicineRecord:		MedicineRecord,
		MedicineQuantity:	TreatmentRecord.MedicineQuantity,
		Treatment:			TreatmentRecord.Treatment,
		Date:				TreatmentRecord.Date,
	}

	//บันทึก
	if err := entity.DB().Create(&lr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lr})

}

// GET: /Treatment/
func ListTreatmentRecord(c *gin.Context) {
	var TreatmentRecord []*entity.TreatmentRecord
	if err := entity.DB().
		Preload("Patient").
		Preload("Doctor").
		Preload("DiagnosisRecord").
		Preload("MedicineRecord").
		Find(&TreatmentRecord).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": TreatmentRecord})
}