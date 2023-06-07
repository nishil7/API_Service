package main

import (
	"github.com/gin-gonic/gin"
	"hms/internal/boot"
	"log"
)

func main() {
	r := gin.Default()
	boot.DatabaseInit()
	boot.Routing(r)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal("Couldn't Able Connect to LocalHost 8080")
	}
}
