package endpoints

import (
	. "github.com/Prabandham/trip_manager/models"
	"github.com/kataras/iris"

	"fmt"
	"math/rand"
	"strings"
	"time"
)

type AuthMiddleware struct{}

func (auth *AuthMiddleware) Serve(c *iris.Context) {
	var person Person
	person.PhoneNumber = c.RequestHeader("Authorization")
	dbperson, err := person.Validate()

	if err == nil && dbperson.PhoneNumber != "" {
		//If user is found set the user and send to next request
		c.Set("current_user", dbperson)
		c.Next()
		return
	}
	// IF user is not found then return error
	response := map[string]string{"message": "No user found"}
	c.JSON(iris.StatusBadRequest, response)
}

// This will be responsible for registering new users.
func RegisterUser(c *iris.Context) {
	var person Person

	if err := c.ReadJSON(&person); err != nil {
		status, response := iris.StatusBadRequest, map[string]string{"message": "Bad request"}
		c.JSON(status, response)
		return
	}

	// We try to find if user exists and if yes we return the existing code.
	dbUser, err := person.Validate()
	if dbUser.PhoneNumber != "" && err == nil {
		status, response := iris.StatusOK, map[string]string{"message": "User exists", "code": dbUser.ConfirmationCode}
		c.JSON(status, response)
		return
	} else {
		person.ConfirmationCode = GenerateCode()
		_, err := person.Save()
		if err == nil {
			status, response := iris.StatusOK, map[string]string{"message": "User saved", "code": person.ConfirmationCode}
			c.JSON(status, response)
			return
		} else {
			status, response := iris.StatusOK, map[string]string{"message": "User not saved"}
			c.JSON(status, response)
			return
		}
	}
}

// This will confirm if the user is present in the db and then validate
// If request code is same as the one generated for them
func ConfirmUser(c *iris.Context) {
	var person Person

	if err := c.ReadJSON(&person); err != nil {
		status, response := iris.StatusBadRequest, map[string]string{"message": "Bad request"}
		c.JSON(status, response)
		return
	}

	dbUser, _ := person.Validate()
	if dbUser.PhoneNumber == "" {
		status, response := iris.StatusBadRequest, map[string]string{"message": "User not found !"}
		c.JSON(status, response)
		return
	}

	dbUser.Confirmed = true
	_, err := dbUser.Save()
	if err == nil {
		status, response := iris.StatusOK, map[string]string{"message": "Verified User"}
		c.JSON(status, response)
		return
	}
}

// This will send out the sms for the user.
func send_sms_code() {
}

// This will generate a random 4 digit number to validate the user
func GenerateCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := arrayToString(r.Perm(5), "")
	return numbers
}

// This will convert an array of ints to a string
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
