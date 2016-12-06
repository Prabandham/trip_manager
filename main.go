package main

import (
	. "github.com/Prabandham/trip_manager/controllers"
	DB "github.com/Prabandham/trip_manager/db"
	//. "github.com/Prabandham/trip_manager/models"

	"gopkg.in/gin-gonic/gin.v1"
	//"fmt"
)

func main() {
	db := DB.Initialize()
	defer db.Close()

	//All this has to move to config or seed file
	//db := DB.Connection()
	//defer db.Close()

	//The below things should run only once and is just for testing.
	//db.DropTableIfExists(&Person{}, &Trip{}, &Expense{})
	//db.AutoMigrate(&Person{}, &Trip{}, &Expense{})

	//DB indexes
	//db.Model(&Person{}).AddUniqueIndex("idx_user_name", "name")
	//db.Model(&Person{}).AddUniqueIndex("idx_phone_number", "phone_number")
	//db.Model(&Person{}).AddIndex("idx_user_name_phone_number", "name", "phone_number")
	//db.Model(&Trip{}).AddIndex("idx_name", "name")
	//db.Model(&Trip{}).AddIndex("idx_location", "location")
	//db.Model(&Trip{}).AddIndex("idx_start_date", "start_date")
	//db.Model(&Expense{}).AddIndex("idx_trip_id", "trip_id")
	//db.Model(&Expense{}).AddIndex("idx_spent", "spent")
	//db.Model(&Expense{}).AddIndex("idx_label", "label")

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
	r.POST("/signup", LoginHandler)

	//All these are application specific routes.
	app := r.Group("/")
	app.Use(AuthRequired())
	{
		app.GET("/trips", TripsIndex)
	}
}
