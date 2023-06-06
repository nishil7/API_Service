package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const charSet string = "abcdefghijklmnopqrstuvwxyz0123456789"

// GenerateUUID : Generates A UUID of Length 5
func GenerateUUID() string {
	currTime := time.Now().UnixNano()
	var uuid string
	size := 5
	for i := 0; i < size; i++ {
		index := currTime % int64(len(charSet))
		uuid += string(charSet[index])
		currTime /= int64(len(charSet))
	}
	return uuid
}

// ErrorHandle : This Function Provides Error Handling of Different Errors
func ErrorHandle(c *gin.Context, err error) {
	//if errors.Is(err, errors.New("record not found")) ||
	//	errors.Is(err, errors.New("ERROR: insert or update on table \\\"patients\\\" violates foreign key constraint \\\"fk_doctors_patient\\\" (SQLSTATE 23503)")) {
	//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	//} else if errors.Is(err, errors.New("ContactNo: invalid length")) ||
	//	errors.Is(err, errors.New("Name: less than min")) ||
	//	errors.Is(err, errors.New("Address: less than min")) ||
	//	errors.Is(err, errors.New("invalid character '\\\"' after object key:value pair")) ||
	//	errors.Is(err, errors.New("invalid character '}' looking for beginning of object key string")) {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	//} else {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	//}
	if errors.Is(err, errors.New("record not found")) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
}
