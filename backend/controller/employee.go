package controller

import (
	"net/http"
	"github.com/AKiRA3563/se-65/entity"
	"github.com/gin-gonic/gin"
	//"github.com/asaskevich/govalidator"
)
// GET /employees
func ListEmployees(c *gin.Context) {
	var employees []entity.Employee
	if err := entity.DB().Raw("SELECT * FROM employees").
		//Preload("Admin").
		Preload("Title").Preload("Role").Preload("Gender").
		Find(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// GET: /employee/:id
func GetEmployee(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM employees WHERE id = ?", id).
		// Preload("Admin").
		Preload("Title").Preload("Role").Preload("Gender").
		Find(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// DELETE /employees/:id
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM employees WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /employees
func UpdateEmployee(c *gin.Context) {
	var employee entity.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if tx := entity.DB().Where("id = ?", employee.ID).First(&employee); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		   return
	}

	if err := entity.DB().Save(&employee).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}