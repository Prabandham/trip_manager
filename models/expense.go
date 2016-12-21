package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	Label       string `db:"label"`
	Description string `db:"description"`
	Spent       int    `db:"spent"`
	Trip        Trip
	TripId      int
	Person      Person
	PersonId    int
}
