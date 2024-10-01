package main

import (
	"todogorest/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.Pinggame)
	r.Run(":" + "5000")
}
