package controllers

import (
	. "github.com/Prabandham/trip_manager/models"
	"gopkg.in/gin-gonic/gin.v1"
	"strings"
)

//This is the middleware that will check for login.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		phone_number := strings.TrimSpace(c.Request.Header.Get("Authorization"))
		//If phone number (Authorization) is not present then
		//it is a bad request or cannot auth user - Skip
		if phone_number != "" {
			person := Person{PhoneNumber: phone_number}
			auth_person, found := person.IsValid()
			if found {
				c.Set("user", auth_person)
				c.Next()
			} else {
				c.AbortWithStatus(403)
			}
		} else {
			c.AbortWithStatus(403)
		}
	}
}

func LoginHandler(c *gin.Context) {
	var person Person
	response := make(map[string]interface{})
	response["message"] = "Invalid User."

	if c.BindJSON(&person) == nil {
		_, err := person.Save()
		if err == nil {
			response["message"] = "User Created Successfully."
		}
	}
	c.JSON(200, response)
}
