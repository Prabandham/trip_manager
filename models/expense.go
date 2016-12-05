package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Expense struct {
	gorm.Model
	Trip        Trip
	TripId      int
	Spent       int
	Description string
	Label       string
}
