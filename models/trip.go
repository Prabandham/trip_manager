package models

import (
	DB "github.com/Prabandham/trip_manager/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"time"
)

type Trip struct {
	gorm.Model
	Name      string
	StartDate time.Time
	EndDate   *time.Time
	Location  string
	People    []Person `gorm:"many2many:person_trips;"`
}

func (trip *Trip) AssociatedPeople() *[]Person {
	db := DB.Connection.New()

	var people []Person
	db.Model(&trip).Related(&people, "People")
	return &people
}

func (trip *Trip) TotalExpense() *int {
	db := DB.Connection.New()

	var totalExpense []int
	db.Raw("select sum(spent) as sum from expenses where trip_id = ?", trip.ID).Pluck("sum", &totalExpense)
	total := totalExpense[0]
	return &total
}
