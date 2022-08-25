package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kackerx/learngin/controllers"
)

func main() {
	server := gin.Default()

	controller := controllers.NewController()
	server.GET("/", controller.GetAll)

	server.Run(":9999")
}
