package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AKiRA3563/se-65/entity"
)

func ListPatient(c *gin.Context) {
	var patient []entity.PatientRegister
	if err := entity.DB().Table("patients").Find(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

func GetPatient(c *gin.Context) {
	id := c.Param("id")
	var patient entity.PatientRegister
	if err := entity.DB().Table("patients").Where("id = ?", id).Find(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}