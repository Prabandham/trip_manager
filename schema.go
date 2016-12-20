// This is going to be a one time call.
// To setup database and tables and indexes.
package main

import (
	"github.com/Prabandham/trip_manager/db"
	. "github.com/Prabandham/trip_manager/models"
)

func LoadSchema() {
	db.Connection.MustExec(`set foreign_key_checks=0;`)

	db.Connection.MustExec("drop table if exists people;")
	db.Connection.MustExec("drop table if exists trips;")
	db.Connection.MustExec("drop table if exists trip_people;")
	db.Connection.MustExec("drop table if exists expenses;")
	db.Connection.MustExec("drop table if exists moments;")

	db.Connection.MustExec(PersonSchema)
	for _, index := range PersonIndexes {
		db.Connection.MustExec(index)
	}
	db.Connection.MustExec(TripSchema)
	for _, index := range TripIndexes {
		db.Connection.MustExec(index)
	}
	db.Connection.MustExec(TripPeopleSchema)
	for _, index := range TripPeopleIndexes {
		db.Connection.MustExec(index)
	}
	db.Connection.MustExec(ExpenseSchema)
	for _, index := range ExpenseIndexes {
		db.Connection.MustExec(index)
	}
	db.Connection.MustExec(MomentsSchema)
	for _, index := range MomentsIndexes {
		db.Connection.MustExec(index)
	}
}
