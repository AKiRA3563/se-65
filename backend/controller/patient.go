package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AKiRA3563/se-65/entity"
)

// GET /patient/:id
func GetPatient(c *gin.Context) {
	var patient entity.PatientRegister
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patient_registers WHERE id = ?", id).
		Preload("Gender").Find(&patient).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET /patients
func ListPatients(c *gin.Context) {
	var patients []entity.PatientRegister
	if err := entity.DB().Raw("SELECT * FROM patient_registers").
		Preload("Gender").Find(&patients).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	c.JSON(http.StatusOK, gin.H{"data": patients})
}