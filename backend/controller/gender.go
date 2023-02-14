package controller

import (
	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGender(c *gin.Context) {
	var gender entity.Gender
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM genders WHERE id = ?", id).Find(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":gender})
}

func ListGenders(c *gin.Context) {
	var gender []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM genders").Find(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gender})
}

