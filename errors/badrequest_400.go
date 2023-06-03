package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Badrequest(err error, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
}
