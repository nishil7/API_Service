package controllers

import (
	"github.com/gin-gonic/gin"
	"hms/models"
	"net/http"
)

func Read_doctor(c *gin.Context) {
	var doctor models.Doctor
	err := models.DB.Where("id = ?", c.Param("id")).Take(&doctor).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Response": doctor})
}
