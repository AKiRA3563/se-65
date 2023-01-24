package controller

import (
	"net/http"

	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
)

// POST	/Diagnosisrecord
func CreateDiagnosisRecord(c *gin.Context) {
	
	var Patient			entity.PatientRegistered
	var Employee		entity.Employee
	var HistorySheet	entity.HistorySheet
	var Disease			entity.Disease
	var DiagnosisRecord entity.DiagnosisRecord

	if err := c.ShouldBindJSON(&DiagnosisRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	//ค้นหา Doctor ด้วย id
	if tx := entity.DB().Where("id = ?", DiagnosisRecord.DoctorID).First(&Employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	//ค้นหา Patient ด้วย id
	if tx := entity.DB().Where("id = ?", DiagnosisRecord.PatientID).First(&Patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	//ค้นหา HistorySheet ด้วย id
	if tx := entity.DB().Where("id = ?", DiagnosisRecord.DoctorID).First(&HistorySheet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "historysheeet not found"})
		return
	}

	//ค้นหา Disease ด้วย id
	if tx := entity.DB().Where("id = ?", DiagnosisRecord.DiseaseID).First(&Disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	//สร้าง DiagnosisRecord
	lr := entity.DiagnosisRecord {
		Patient: 		Patient,
		Doctor:  		Employee,
		HistorySheet: 	HistorySheet,
		Disease: 		Disease,
		Examination: 	DiagnosisRecord.Examination,
		MedicalCertificate: DiagnosisRecord.MedicalCertificate,
		Date:			DiagnosisRecord.Date,
	}

	//บันทึก
	if err := entity.DB().Create(&lr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lr})

}

// GET: /Diagnosis/
func ListDiagnosisRecord(c *gin.Context) {
	var DiagnosisRecord []*entity.DiagnosisRecord
	if err := entity.DB().
		Preload("Patient").
		Preload("Doctor").
		Preload("HistorySheet").
		Preload("Disease").
		Find(&DiagnosisRecord).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": DiagnosisRecord})
}
