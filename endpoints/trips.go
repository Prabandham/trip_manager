package endpoints

import (
	"github.com/kataras/iris"
)

// This will return the current trip details
func CurrentTrip(c *iris.Context) {
	c.JSON(iris.StatusOK, map[string]string{"message": "You will see this only if authenticated"})
}
