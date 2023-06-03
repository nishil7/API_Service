package main

import (
	"github.com/gin-gonic/gin"
	"hms/models"
	"hms/router"
)

func main() {
	r := gin.Default()
	models.Database_connect()
	router.Routing(r)
	r.Run("localhost:8080")
}
