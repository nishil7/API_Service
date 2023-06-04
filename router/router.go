package router

import (
	"github.com/gin-gonic/gin"
	"hms/controllers"
)

func Routing(r *gin.Engine) {
	r.POST("/doctor/", controllers.Create_doctor)
	r.GET("/doctor/:id", controllers.Read_doctor)
	r.PATCH("/doctor/:id", controllers.Update_doctor)
	r.POST("/patient/", controllers.Create_patient)
	r.GET("/patient/", controllers.Read_patient)
	r.PATCH("/patient/", controllers.Update_patient)
	r.GET("/fetchPatientByDoctorId/:id", controllers.Fetch_patients_by_doctor)
}
