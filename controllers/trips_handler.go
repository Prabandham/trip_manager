package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func TripsIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Authenticated successfully !!!",
	})
}
