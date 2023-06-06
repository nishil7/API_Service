package doctor

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

// NewController : This function initialize the Controller.
func NewController(core ICore) *Controller {
	return &Controller{
		core: core,
	}
}

// FetchById  : This function fetches the doctor id parameter and provides with the response.
func (c *Controller) FetchById(ctx *gin.Context) {
	id := ctx.Param("id")
	doctorResponse, err := c.core.FetchById(id)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": doctorResponse})
}

// Create : This function requests for the creation of the new doctor details.
func (c *Controller) Create(ctx *gin.Context) {
	var data _struct.NewDoctorRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	doctorResponse, err := c.core.Create(&data)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": doctorResponse})
}

// Update : This function passes the update details request of doctor and responds with the updated details.
func (c *Controller) Update(ctx *gin.Context) {
	var data _struct.UpdateDoctorRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ErrorHandle(ctx, err)
		log.Printf(err.Error())
		return
	}
	id := ctx.Param("id")
	doctorResponse, err := c.core.Update(&data, id)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": doctorResponse})
}

// FetchPatientByDoctorId : This function fetches the patient with corresponding doctor ids
func (c *Controller) FetchPatientByDoctorId(ctx *gin.Context) {
	id := ctx.Param("id")
	patientResponse, doctorResponse, err := c.core.FetchPatientByDoctorId(id)
	if err != nil {
		common.ErrorHandle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": patientResponse, "Doctor": doctorResponse})
}
