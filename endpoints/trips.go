package endpoints

import (
	. "github.com/Prabandham/trip_manager/models"
	"github.com/kataras/iris"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"fmt"
)

// This will return the current trip details
func CurrentTrip(c *iris.Context) {
	person := c.Get("current_user").(*Person)
	var trips []Trip
	db, _ := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close()

	db.Model(&person).Related(&trips, "Trips")
	c.JSON(iris.StatusOK, &trips)
}

func CreateTrip(c *iris.Context) {
	person := c.Get("current_user").(*Person)
	var trip Trip

	if err := c.ReadJSON(&trip); err != nil {
		status, response := iris.StatusBadRequest, map[string]string{"message": "Bad request"}
		c.JSON(status, response)
		return
	}

	trip.People = []Person{*person}
	db, _ := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close()

	db.Save(&trip)
	c.JSON(iris.StatusOK, map[string]interface{}{"payload": &trip})
}

// Add people to trip
func AddBuddies(c *iris.Context) {
	// If current User is not there in the trip then he cannot add another person only.
	type params struct {
		trip_id int      `json:"trip_id"`
		buddies []Person `json:"buddies"`
	}
	var p params
	var trip Trip

	err := c.ReadJSON(&p)

	if err != nil {
		status, response := iris.StatusBadRequest, map[string]string{"message": "Bad request"}
		c.JSON(status, response)
		return
	}

	db, _ := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close()

	fmt.Println(p.buddies)
	fmt.Println(p.trip_id)

	for _, person := range p.buddies {
		var dbUser Person
		// Try to find if User exist on system
		db.Where("phone_number = ?", person.PhoneNumber).Find(&dbUser)
		if dbUser.PhoneNumber == "" {
			//If Not We will create a User
			db.Create(&person)
		}
		db.First(&trip, p.trip_id)
		trip.People = []Person{person}
		db.Save(&trip)
	}
	c.JSON(iris.StatusOK, map[string]interface{}{"message": "Created buddies", "payload": &trip})
}
