package controllers

import (
	"github.com/gin-gonic/gin"
	"hms/errors"
	"hms/models"
	"net/http"
)

type update_doctor_info struct {
	Contact_No string `json:"contact_no"`
}

func Update_doctor(c *gin.Context) {
	var doctor models.Doctor
	err := models.DB.Where("id = ?", c.Query("id")).Take(&doctor).Error
	if err != nil {
		errors.Statusnot_404(err, c)
		return
	}
	var data update_doctor_info
	err = c.ShouldBindJSON(&data)
	if err != nil {
		errors.Badrequest(err, c)
		return
	}
	updated_doctor := models.Doctor{Contact_no: data.Contact_No}
	models.DB.Model(&doctor).Updates(&updated_doctor)
	c.JSON(http.StatusOK, gin.H{"Response": doctor})

}
