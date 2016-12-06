package controllers

import (
	"fmt"
	. "github.com/Prabandham/trip_manager/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func TripsIndex(c *gin.Context) {
	current_user := c.MustGet("user").(*Person)
	fmt.Println(current_user)
	trips := current_user.Trips()
	c.JSON(200, &trips)
}
