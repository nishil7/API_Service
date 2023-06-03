package controllers

import (
	"github.com/gin-gonic/gin"
	"hms/errors"
	"hms/models"
	"net/http"
)

type create_patient_info struct {
	Name       string `json:"name"`
	Contact_no string `json:"contact_no"`
	Address    string `json:"address"`
	Doctor_id  int    `json:"doctor_id"`
}

func Create_patient(c *gin.Context) {
	var data create_patient_info
	err := c.ShouldBindJSON(&data)
	if err != nil {
		errors.Badrequest(err, c)
		return
	}
	patient := models.Patient{Name: data.Name, Contact_no: data.Contact_no, Address: data.Address, Doctor_id: data.Doctor_id}
	err = models.DB.Create(&patient).Error
	if err != nil {
		errors.Badrequest(err, c)
		return
	}
	models.DB.Where("id = ?", patient.Doctor_id).Take(&patient.Doctor)
	c.JSON(http.StatusOK, gin.H{"Response": patient})
}
