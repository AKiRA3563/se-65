package controller

import (
	"net/http"

	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
)

// POST	/medicalcertificates
func CreateMedicalCertificate(c *gin.Context) {
	var medicalcertificate	entity.MedicalCertificate
	if err := c.ShouldBindJSON(&medicalcertificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&medicalcertificate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicalcertificate})
}

// GET /medicalcertificate/:id
func GetMedicalCertificate(c *gin.Context) {
	var medicalcertificate	entity.MedicalCertificate
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medical_certificates WHERE id = ?", id).
		Find(&medicalcertificate).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicalcertificate})
}

// GET /medicalcertificates
func ListMedicalCertificates(c *gin.Context) {
	var medicalcertificates	[]entity.MedicalCertificate
	if err := entity.DB().Raw("SELECT * FROM medical_certificates").
		Find(&medicalcertificates).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicalcertificates})
}