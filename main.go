package main

import (
	. "github.com/Prabandham/trip_manager/controllers"
	DB "github.com/Prabandham/trip_manager/db"
	. "github.com/Prabandham/trip_manager/models"

	"gopkg.in/gin-gonic/gin.v1"
	//"fmt"
)

func main() {
	db := DB.Connection()
	defer db.Close()

	//The below things should run only once and is just for testing.
	//db.DropTableIfExists(&Person{}, &Trip{}, &Expense{})
	db.AutoMigrate(&Person{}, &Trip{}, &Expense{})

	//Seed data.
	//SeedPeople()
	//CreateTrip()
	//AddExpense()

	//gin.SetMode(gin.ReleaseMode) This will set production mode
	r := gin.Default()
	LoadRoutes(r)

	r.Run()
}

func LoadRoutes(r *gin.Engine) {
	//This is just to test if server is up and running.
	r.GET("/ping", PingHandler)
	r.POST("/login", LoginHandler)

	//All these are application specific routes.
	app := r.Group("/")
	app.Use(AuthRequired)
	{
		app.GET("/trips", TripsIndex)
	}
}
