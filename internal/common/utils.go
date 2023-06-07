package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

const charSet string = "abcdefghijklmnopqrstuvwxyz0123456789"
const size = 5

// GenerateUUID : Generates A UUID of Length 5
func GenerateUUID() string {
	currTime := time.Now().UnixNano()
	var uuid string
	for i := 0; i < size; i++ {
		index := currTime % int64(len(charSet))
		uuid += string(charSet[index])
		currTime /= int64(len(charSet))
	}
	return uuid
}

func RecordNotFound(err error) bool {
	return gorm.ErrRecordNotFound == err
}
func Validator(err error) bool {
	return strings.Contains(err.Error(), validator.ErrLen.Error()) || strings.Contains(err.Error(), validator.ErrInvalid.Error()) ||
		strings.Contains(err.Error(), validator.ErrRegexp.Error()) || strings.Contains(err.Error(), validator.ErrMin.Error())
}

//func ForiegnKey(err error) bool {
//	return gorm.ErrForeignKeyViolated == err
//}

// ErrorHandle : This Function Provides Error Handling of Different Errors
func ErrorHandle(c *gin.Context, err error) {
	if RecordNotFound(err) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	} else if Validator(err) || c.Err() == err {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
}
