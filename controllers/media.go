package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetMediaFile - Finds media file for the requested parameter
func GetMediaFile(c *gin.Context) {
	file := c.Param("file")

	c.File("resources/media/" + file)
}
