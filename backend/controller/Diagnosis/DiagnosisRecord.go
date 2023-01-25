package controller

import (
	"net/http"

	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
)

// POST	/diagnosis_records
func CreateDiagnosisRecord(c *gin.Context) {
	
	var patient			entity.PatientRegister
	var employee		entity.Employee
	var historysheet	entity.HistorySheet
	var disease			entity.Disease
	var diagnosisrecord entity.DiagnosisRecord

	if err := c.ShouldBindJSON(&diagnosisrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	//ค้นหา Doctor ด้วย id
	if tx := entity.DB().Where("id = ?", diagnosisrecord.DoctorID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	//ค้นหา Patient ด้วย id
	if tx := entity.DB().Where("id = ?", diagnosisrecord.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	//ค้นหา HistorySheet ด้วย id
	if tx := entity.DB().Where("id = ?", diagnosisrecord.DoctorID).First(&historysheet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "historysheeet not found"})
		return
	}

	//ค้นหา Disease ด้วย id
	if tx := entity.DB().Where("id = ?", diagnosisrecord.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	//สร้าง DiagnosisRecord
		dr := entity.DiagnosisRecord {
			Patient: 		patient,
			Doctor:  		employee,
			HistorySheet: 	historysheet,
			Disease: 		disease,
			Examination: 	diagnosisrecord.Examination,
			MedicalCertificate: diagnosisrecord.MedicalCertificate,
			Date:			diagnosisrecord.Date,
		}

	//บันทึก
	if err := entity.DB().Create(&dr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	c.JSON(http.StatusOK, gin.H{"data": dr})

}

// GET /diagnosisrecord/:id
func GetDiagnosisRecord(c *gin.Context) {
	var diagnosisrecord entity.DiagnosisRecord
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&diagnosisrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diagnosisrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diagnosisrecord})
}

// GET /diagnosis_records/
func ListDiagnosisRecord(c *gin.Context) {
	var diagnosisrecords []*entity.DiagnosisRecord
	if err := entity.DB().
		Preload("patient").
		Preload("doctor").
		Preload("historySheet").
		Preload("disease").
		Find(&diagnosisrecords).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": diagnosisrecords})
}

// DELETE /diagnosis_records/:id
func DeleteDiagnosisRecord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM diagnosis_records WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diagnosisrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

//PATCH /diagnosis_records
func UpdateDiagnosisRecord(c *gin.Context) {
	var diagnosisrecord entity.DiagnosisRecord
	if err := c.ShouldBindJSON(&diagnosisrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", diagnosisrecord.ID).First(&diagnosisrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diagnosis not found"})
		return
	}

	if err := entity.DB().Save(&diagnosisrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diagnosisrecord})
}