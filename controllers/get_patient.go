package controllers

import (
	"github.com/gin-gonic/gin"
	"hms/errors"
	"hms/models"
	"net/http"
)

func Read_patient(c *gin.Context) {
	var patient models.Patient
	err := models.DB.Where("id = ?", c.Query("id")).Take(&patient).Error
	if err != nil {
		errors.Statusnot_404(err, c)
		return
	}
	models.DB.Where("id = ?", patient.Doctor_id).Take(&patient.Doctor)
	c.JSON(http.StatusOK, gin.H{"Response": patient})
}
