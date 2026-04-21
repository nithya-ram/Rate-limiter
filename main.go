package main

import (
	"rate-limiter/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/request", controllers.HandleRequest)

	r.GET("/status", controllers.GetStatus)

	r.Run(":8080")
}
