package controllers

import (
	"github.com/gin-gonic/gin"
	"hms/errors"
	"hms/models"
	"net/http"
)

type update_patient_info struct {
	Doctor_id  int    `json:"doctor_id"`
	Contact_No string `json:"contact_no"`
	Address    string `json:"address"`
}

func Update_patient(c *gin.Context) {
	var patient models.Patient
	err := models.DB.Where("id = ?", c.Query("id")).Take(&patient).Error
	if err != nil {
		errors.Statusnot_404(err, c)
		return
	}
	var data update_patient_info
	err = c.ShouldBindJSON(&data)
	if err != nil {
		errors.Badrequest(err, c)
		return
	}
	updated_patient := models.Patient{Contact_no: data.Contact_No, Address: data.Address, Doctor_id: data.Doctor_id}
	err = models.DB.Model(&patient).Updates(&updated_patient).Error
	if err != nil {
		errors.Badrequest(err, c)
		return
	}
	models.DB.Where("id = ?", patient.Doctor_id).Take(&patient.Doctor)
	c.JSON(http.StatusOK, gin.H{"Response": patient})

}
