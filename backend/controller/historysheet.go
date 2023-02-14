package controller

import (
	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /historysheet/:id
func GetHistoySheet(c *gin.Context) {
	var historysheet entity.HistorySheet
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM history_sheets WHERE id = ?", id).
		Preload("PatientRegister").
		Preload("Nurse").Find(&historysheet).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data":historysheet})
}


// GET /historysheets
func ListHistorysheets(c *gin.Context) {
	var historysheet []entity.HistorySheet
	if err := entity.DB().Raw("SELECT * FROM history_sheets").
		Preload("PatientRegister").
		Preload("Nurse").Find(&historysheet).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"data": historysheet})
}
