package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Up and Running !!",
	})
}
