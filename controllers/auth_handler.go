package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gin-gonic/gin.v1"
)

//Global struct that defines JWT data
type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

//This is the middleware that will check for login.
func AuthRequired(c *gin.Context) {
	if c.Query("Auth") == "Test" {
		//Continue
	} else {
		c.AbortWithStatus(403)
	}
}

func LoginHandler(c *gin.Context) {
	setToken(c)
}

func setToken(c *gin.Context) {

}
