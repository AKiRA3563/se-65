package main

import (

	"github.com/gin-gonic/gin"
	"github.com/AKiRA3563/se-65/entity"
	"github.com/AKiRA3563/se-65/controller"
	//"github.com/AKiRA3563/se-65/middlewares"
)

const PORT = "8080"

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		// router.Use(middlewares.Authorizes())
		// {
			// DiagonsisRecord Routes
			router.GET("/diagnosis_records", controller.ListDiagnosisRecord)
			router.GET("/diagnosisrecords/:id", controller.GetDiagnosisRecord)
			router.POST("/diagnosis_record", controller.CreateDiagnosisRecord)
			router.PATCH("/diagnosis_records", controller.UpdateDiagnosisRecord)
			router.DELETE("/diagnosisrecords/:id", controller.DeleteDiagnosisRecord)

			// TreatmentRecord Routes
			router.GET("/treatment_records", controller.ListTreatmentRecord)
			router.GET("/treatmentrecord/:id", controller.GetTreatmentRecord)
			router.POST("/treatment_records", controller.CreateTreatmentRecord)
			router.PATCH("/treatment_records", controller.UpdateTreatmentRecord)
			router.DELETE("/treatmentrecords/:id", controller.DeleteTreatmentRecord)

		// }
	}

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
