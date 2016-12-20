package endpoints

import (
	"github.com/kataras/iris"

	"github.com/Prabandham/trip_manager/db"
	. "github.com/Prabandham/trip_manager/models"

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
	var status int
	var response = make(map[string]string)

	if err := c.ReadJSON(&person); err != nil {
		status, response = iris.StatusBadRequest, map[string]string{"message": "Bad request"}
	}

	// We try to find if user exists and if yes we return the existing code.
	dbUser, err := person.Validate()
	if dbUser.PhoneNumber != "" && err == nil {
		status, response = iris.StatusOK, map[string]string{"message": "User exists", "code": dbUser.ConfirmationCode}
	} else {
		// We create the user.
		person.CreatedAt = time.Now().Format(time.RFC3339)
		person.ConfirmationCode = generate_code()

		db := db.Connection
		tx := db.MustBegin()
		tx.NamedExec("INSERT INTO people (user_name,phone_number,confirmation_code,created_at) VALUES (:user_name,:phone_number,:confirmation_code,:created_at)", &person)
		tx.Commit()

		status, response = iris.StatusOK, map[string]string{"message": "User saved", "code": person.ConfirmationCode}
	}
	c.JSON(status, response)
}

// This will confirm if the user is present in the db and then validate
// If request code is same as the one generated for them
func ConfirmUser(c *iris.Context) {
	var person Person
	var status int
	var response = make(map[string]string)

	if err := c.ReadJSON(&person); err != nil {
		status, response = iris.StatusBadRequest, map[string]string{"message": "Bad request"}
	}

	dbUser, _ := person.Validate()
	if dbUser.PhoneNumber == "" {
		status, response = iris.StatusBadRequest, map[string]string{"message": "User not found !"}
		c.JSON(status, response)
		return
	}

	res := db.Connection.MustExec("UPDATE people SET confirmed=? WHERE id=?", "YES", dbUser.Id)
	affect, _ := res.RowsAffected()

	if affect != 1 {
		status, response = iris.StatusForbidden, map[string]string{"message": "User already confirmed !"}
		c.JSON(status, response)
		return
	}
	status, response = iris.StatusOK, map[string]string{"message": "Verified User"}
	c.JSON(status, response)
}

// This should get called on Uninstall of the application.
func DeleteUser(c *iris.Context) {
	var person Person
	var status int
	var response = make(map[string]string)

	if err := c.ReadJSON(&person); err != nil {
		status, response = iris.StatusBadRequest, map[string]string{"message": "Bad request"}
	}
	res := db.Connection.MustExec("DELETE FROM people where id=?", person.Id)
	count, _ := res.RowsAffected()
	if count != 1 {
		status, response = iris.StatusBadRequest, map[string]string{"message": "User Not found"}
	}
	status, response = iris.StatusOK, map[string]string{"message": "Thank you !"}
	c.JSON(status, response)
}

// This will send out the sms for the user.
func send_sms_code() {
}

// This will generate a random 4 digit number to validate the user
func generate_code() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := arrayToString(r.Perm(5), "")
	return numbers
}

// This will convert an array of ints to a string
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
