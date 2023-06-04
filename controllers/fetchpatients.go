package controllers

import (
	"github.com/gin-gonic/gin"
	"hms/errors"
	"hms/models"
	"net/http"
)

func Fetch_patients_by_doctor(c *gin.Context) {
	var patients []models.Patient
	err := models.DB.Where("doctor_id = ?", c.Param("id")).Find(&patients).Error
	if err != nil {
		errors.Statusnot_404(err, c)
		return
	}
	for i := range patients {
		models.DB.Where("id = ?", c.Param("id")).Take(&patients[i].Doctor)
	}
	c.JSON(http.StatusOK, gin.H{"Response": patients})
}
