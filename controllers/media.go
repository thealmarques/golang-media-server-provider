package controllers

import (
	"github.com/gin-gonic/gin"
)

// FindMedia - Finds media file for the requested parameter
func FindMedia(c *gin.Context) {
	file := c.Param("file")

	c.JSON(200, gin.H{
		"filename": file,
	})
}
