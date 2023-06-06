package main

import (
	"github.com/gin-gonic/gin"
	"hms/internal/boot"
)

func main() {
	r := gin.Default()
	boot.DatabaseInit()
	boot.Routing(r)
	r.Run("localhost:8080")
}
