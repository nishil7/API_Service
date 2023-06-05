package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"hms/errors"
	"hms/models"
	"net/http"
)

type create_patient_info struct {
	Name       string `json:"name" validate:"min = 1"`
	Contact_no string `json:"contact_no" validate:"len = 10,regexp = ^[0-9]+$"`
	Address    string `json:"address" validate:"min = 1"`
	Doctor_id  int    `json:"doctor_id"`
}

func Create_patient(c *gin.Context) {
	var data create_patient_info
	err := c.ShouldBindJSON(&data)
	if err != nil {
		errors.Badrequest(err, c)
		return
	}
	err = validator.Validate(data)
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
