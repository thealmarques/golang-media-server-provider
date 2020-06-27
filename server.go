package main

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/media/:file", controllers.GetMediaFile)

	router.Run("0.0.0.0:9000")
}
