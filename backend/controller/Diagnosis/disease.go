package controller

import (
	"net/http"

	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
)

// POST	/diseases
func CreateDisease(c *gin.Context) {
	var disease	entity.Disease
	if err := entity.DB().Create(&disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	c.JSON(http.StatusOK, gin.H{"data": disease})

}

// GET /disease/:id
func GetDisease(c *gin.Context) {
	var disease	entity.Disease
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diagnosisrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": disease})
}

// GET /diseases/
func ListDisease(c *gin.Context) {
	var diseases	entity.Disease
	if err := entity.DB().Raw("SELECT * FROM diseases").Find (&diseases).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": diseases})
}

// DELETE /diseases/:id
func DeleteDisease(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM diiseases WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

//PATCH /diseases
func UpdateDisease(c *gin.Context) {
	var disease	entity.Disease
	if err := c.ShouldBindJSON(&disease); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", disease.ID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	if err := entity.DB().Save(&disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": disease})
}