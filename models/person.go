package models

import (
	DB "github.com/Prabandham/trip_manager/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	gorm.Model
	Name        string
	PhoneNumber string `gorm:"not null;unique"`
}

func (person *Person) Trips() *[]Trip {
	db := DB.Connection()
	defer db.Close()

	var person_trip_ids []int
	var trips []Trip

	db.Raw("select trip_id from person_trips where person_id = ?", person.ID).Pluck("trip_id", &person_trip_ids)
	db.Where("id in (?)", person_trip_ids).Find(&trips)

	return &trips
}

func (person *Person) ActiveTrip() *Trip {
	db := DB.Connection()
	defer db.Close()

	var trip_ids []int
	var trip Trip

	db.Raw("select id from trips as t left join person_trips as pt on pt.trip_id=t.id where pt.person_id = ?", person.ID).Pluck("id", &trip_ids)
	db.Where("id in (?)", trip_ids).Last(&trip)

	return &trip
}
