package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"hms/errors"
	"hms/models"
	"net/http"
)

type Create_doctor_info struct {
	Name       string `json:"name" validate:"min = 1"`
	Contact_no string `json:"contact_no" validate:"len = 10,regexp = ^[0-9]+$"`
}

func Create_doctor(c *gin.Context) {
	var data Create_doctor_info
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
	doctor := models.Doctor{Name: data.Name, Contact_no: data.Contact_no}
	models.DB.Create(&doctor)
	c.JSON(http.StatusOK, gin.H{"Response": doctor})
}
