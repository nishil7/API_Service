package patient

import (
	"github.com/gin-gonic/gin"
	"hms/internal/common"
	"hms/internal/common/struct"
	"log"
	"net/http"
)

type Controller struct {
	core ICore
}

// NewController : This function initializes the controller of the patient.
func NewController(core ICore) *Controller {
	return &Controller{
		core: core,
	}
}

// FetchById : This function fetches the patient id parameter and provides with the response.
func (c *Controller) FetchById(ctx *gin.Context) {
	id := ctx.Param("id")
	patientResponse, err := c.core.FetchById(id)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Response": patientResponse})
}

// Create : This function requests for the creation of the new patient details.
func (c *Controller) Create(ctx *gin.Context) {
	var data _struct.NewPatientRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ErrorHandle(ctx, err)
		log.Printf(err.Error())
		return
	}
	doctorResponse, err := c.core.Create(&data)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Response": doctorResponse})
}

// Update : This function passes the update details request of patient and responds with the updated details.
func (c *Controller) Update(ctx *gin.Context) {
	var data _struct.UpdatePatientRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ErrorHandle(ctx, err)
		log.Printf(err.Error())
		return
	}
	id := ctx.Param("id")
	patientResponse, err := c.core.Update(&data, id)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Response": patientResponse})
}
