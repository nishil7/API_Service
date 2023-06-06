package boot

import (
	"github.com/gin-gonic/gin"
)

// Routing : Provides with routing of different URLs
func Routing(r *gin.Engine) {
	r.POST("/doctor/", DoctorController.Create)
	r.GET("/doctor/:id", DoctorController.FetchById)
	r.PATCH("/doctor/:id", DoctorController.Update)
	r.POST("/patient/", PatientController.Create)
	r.GET("/patient/:id", PatientController.FetchById)
	r.PATCH("/patient/:id", PatientController.Update)
	r.GET("/fetchPatientByDoctorId/:id", DoctorController.FetchPatientByDoctorId)
}
