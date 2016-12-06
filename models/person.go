package models

import (
	"errors"

	DB "github.com/Prabandham/trip_manager/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `gorm:"not null;unique" json:"phone_number" binding:"required"`
}

func (person *Person) Trips() *[]Trip {
	db := DB.Connection.New()

	var person_trip_ids []int
	var trips []Trip

	db.Raw("select trip_id from person_trips where person_id = ?", person.ID).Pluck("trip_id", &person_trip_ids)
	db.Where("id in (?)", person_trip_ids).Find(&trips)

	return &trips
}

func (person *Person) ActiveTrip() *Trip {
	db := DB.Connection.New()

	var trip_ids []int
	var trip Trip

	db.Raw("select id from trips as t left join person_trips as pt on pt.trip_id=t.id where pt.person_id = ?", person.ID).Pluck("id", &trip_ids)
	db.Where("id in (?)", trip_ids).Last(&trip)

	return &trip
}

func (person *Person) IsValid() (*Person, bool) {
	db := DB.Connection.New()

	var auth_person Person
	if person.PhoneNumber != "" {
		db.Where("phone_number = ?", person.PhoneNumber).First(&auth_person)
		//If the user could not be found then NewRecord will be true
		if db.NewRecord(auth_person) {
			return nil, false
		} else {
			return &auth_person, true
		}
	}
	return nil, false
}

func (person *Person) Save() (*Person, error) {
	db := DB.Connection.New()

	var new_person Person
	db.Where(&person).FirstOrCreate(&new_person)

	if db.NewRecord(new_person) {
		return nil, errors.New("Person Could not be saved !!")
	}
	return &new_person, nil
}
