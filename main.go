package main

import (
	. "github.com/Prabandham/trip_manager/endpoints"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {
	// Initialize iris to use logger
	iris.Use(logger.New())

	//LoadSchema()

	// This is going act as the routes.
	LoadRoutes()

	// Default PORT
	iris.Listen(":3000")
}

func LoadRoutes() {
	iris.Get("/ping", Pong)
	iris.Post("/register", RegisterUser)
	iris.Post("/confirm", ConfirmUser)

	//These are all protected Routes
	iris.Use(&AuthMiddleware{})
	iris.Get("/trips", CurrentTrip)
	iris.Post("/create_trip", CreateTrip)
	iris.Post("/add_buddies", AddBuddies)
}

func Pong(c *iris.Context) {
	c.JSON(iris.StatusOK, map[string]string{"status": "Test !!"})
}
