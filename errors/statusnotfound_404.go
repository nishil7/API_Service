package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Statusnot_404(err error, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
}
