package main

import (
	
	//"github.com/AKiRA3563/se-65/controller"
	controller_diagnosis "github.com/AKiRA3563/se-65/controller/Diagnosis"
	controller_treatment "github.com/AKiRA3563/se-65/controller/Treatment"

	//"github.com/AKiRA3563/se-65/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/AKiRA3563/se-65/entity"
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
			router.GET("/diagnosis_records", controller_diagnosis.ListDiagnosisRecord)
			router.GET("/diagnosisrecords/:id", controller_diagnosis.GetDiagnosisRecord)
			router.POST("/diagnosis_record", controller_diagnosis.CreateDiagnosisRecord)
			router.PATCH("/diagnosis_records", controller_diagnosis.UpdateDiagnosisRecord)
			router.DELETE("/diagnosisrecords/:id", controller_diagnosis.DeleteDiagnosisRecord)

			// TreatmentRecord Routes
			router.GET("/treatment_records", controller_treatment.ListTreatmentRecord)
			router.GET("/treatmentrecord/:id", controller_treatment.GetTreatmentRecord)
			router.POST("/treatment_records", controller_treatment.CreateTreatmentRecord)
			router.PATCH("/treatment_records", controller_treatment.UpdateTreatmentRecord)
			router.DELETE("/treatmentrecords/:id", controller_treatment.DeleteTreatmentRecord)

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
