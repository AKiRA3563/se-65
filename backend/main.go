package main

import (
	"github.com/AKiRA3563/se-65/controller"
	"github.com/AKiRA3563/se-65/entity"
	"github.com/AKiRA3563/se-65/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// DiagonsisRecord Routes
			r.GET("/diagnosisrecords", controller.ListDiagnosisRecords)
			r.GET("/diagnosisrecord/:id", controller.GetDiagnosisRecord)
			r.POST("/diagnosisrecords", controller.CreateDiagnosisRecord)
			r.PATCH("/diagnosisrecords", controller.UpdateDiagnosisRecord)
			r.DELETE("/diagnosisrecords/:id", controller.DeleteDiagnosisRecord)

			// TreatmentRecord Routes
			r.GET("/treatmentrecords", controller.ListTreatmentRecords)
			r.GET("/treatmentrecord/:id", controller.GetTreatmentRecord)
			r.POST("/treatmentrecords", controller.CreateTreatmentRecord)
			r.PATCH("/treatmentrecords", controller.UpdateTreatmentRecord)
			r.DELETE("/treatmentrecords/:id", controller.DeleteTreatmentRecord)

			// Disease Routes
			r.GET("/diseases", controller.ListDiseases)
			r.GET("/disease/:id", controller.GetDisease)
			r.POST("/diseases", controller.CreateDisease)
			r.PATCH("/diseases", controller.UpdateDisease)
			r.DELETE("/diseases/:id", controller.DeleteDisease)

			// Medicine Routes
			r.GET("/medicines", controller.ListMedicines)
			r.GET("/medicine/:id", controller.GetMedicine)
			r.POST("/medicines", controller.CreateMedicine)
			r.PATCH("/medicines", controller.UpdateMedicine)
			r.DELETE("/medicines/:id", controller.DeleteMedicine)

			r.GET("/employees", controller.ListEmployees)
			r.GET("/employee/:id", controller.GetEmployee)

			r.GET("/patients", controller.ListPatients)
			r.GET("/patient/:id", controller.GetPatient)

			r.GET("/historysheets", controller.ListHistorysheets)
			r.GET("/historysheet/:id", controller.GetHistoySheet)

			r.GET("/gender", controller.ListGenders)
			r.GET("gender/:id", controller.GetGender)
		}
	}

	// login User Route
	r.POST("/login", controller.Login)

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
