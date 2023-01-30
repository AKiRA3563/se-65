package controller

import (
	"net/http"

	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
)

// Post /treament_records
func CreateTreatmentRecord(c *gin.Context) {

	var patient			entity.PatientRegister
	var employee		entity.Employee
	var diagnosisrecord entity.DiagnosisRecord
	var medicine		entity.Medicine
	var treatmentrecord	entity.TreatmentRecord

	
	if err := c.ShouldBindJSON(&treatmentrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา Doctor ด้วย id
	if tx := entity.DB().Where("id = ?", treatmentrecord.DoctorID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	//ค้นหา Patient ด้วย id
	if tx := entity.DB().Where("id = ?", treatmentrecord.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	//ค้นหา Medicine ด้วย id
	if tx := entity.DB().Where("id = ?", treatmentrecord.MedicineID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	//สร้าง TreamentRecord
	tr := entity.TreatmentRecord {
		Patient:			patient,
		Doctor:				employee,
		DiagnosisRecord:	diagnosisrecord,
		Medicine:			medicine,
		MedicineQuantity:	treatmentrecord.MedicineQuantity,
		Treatment:			treatmentrecord.Treatment,
		Date:				treatmentrecord.Date,
	}

	//บันทึก
	if err := entity.DB().Create(&tr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tr})

}

// GET /treatmentrecord/:id
func GetTreatmentRecord(c *gin.Context) {
	var treatmentrecord entity.TreatmentRecord
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&treatmentrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatmentrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatmentrecord})
}

// GET /treatment_record/
func ListTreatmentRecord(c *gin.Context) {
	var treatmentrecords []*entity.TreatmentRecord
	if err := entity.DB().
		Preload("patient").
		Preload("doctor").
		Preload("diagnosisrecord").
		Preload("medicine").
		Find(&treatmentrecords).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatmentrecords})
}

// DELETE /treament_records/:id
func DeleteTreatmentRecord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM treatment_records WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatmentrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

//PATCH /diagnosis_records
func UpdateTreatmentRecord(c *gin.Context) {
	var treatmentrecord	entity.TreatmentRecord
	if err := c.ShouldBindJSON(&treatmentrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", treatmentrecord.ID).First(&treatmentrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diagnosis not found"})
		return
	}

	if err := entity.DB().Save(&treatmentrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatmentrecord})
}