package controllers

import (
	"github.com/gin-gonic/gin"
	"hms/errors"
	"hms/models"
	"net/http"
)

type Create_doctor_info struct {
	Name       string `json:"name"`
	Contact_no string `json:"contact_no"`
}

func Create_doctor(c *gin.Context) {
	var data Create_doctor_info
	err := c.ShouldBindJSON(&data)
	if err != nil {
		errors.Badrequest(err, c)
	}
	doctor := models.Doctor{Name: data.Name, Contact_no: data.Contact_no}
	models.DB.Create(&doctor)
	c.JSON(http.StatusOK, gin.H{"Response": doctor})
}
