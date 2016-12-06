package main

import (
	DB "github.com/Prabandham/trip_manager/db"
	. "github.com/Prabandham/trip_manager/models"

	"time"
)

func SeedPeople() {
	db := DB.Connection.New()
	defer db.Close()

	person1 := Person{Name: "Prabandham Srinidhi", PhoneNumber: "9738912733"}
	person2 := Person{Name: "Abhishek Murthy", PhoneNumber: "9986785453"}
	person3 := Person{Name: "Krishna S", PhoneNumber: "9738912816"}
	person4 := Person{Name: "Ashwin J", PhoneNumber: "9738107286"}
	person5 := Person{Name: "Lohith B N", PhoneNumber: "9480114203"}

	db.Create(&person1)
	db.Create(&person2)
	db.Create(&person3)
	db.Create(&person4)
	db.Create(&person5)
}

func CreateTrip() {
	db := DB.Connection.New()
	defer db.Close()

	var people []Person
	db.Find(&people)

	trip := Trip{
		Name:      "Pondycherry",
		StartDate: time.Now(),
		Location:  "Pondycherry",
		People:    people,
	}
	db.Create(&trip)
}

func AddExpense() {
	db := DB.Connection.New()
	defer db.Close()

	var trip Trip
	db.Last(&trip)

	expense1 := Expense{Trip: trip, Spent: 3850, Description: "Bus Tickets", Label: "Travel"}
	expense2 := Expense{Trip: trip, Spent: 500, Description: "While Smoking in bus stop", Label: "Ciggarates and Drinks"}

	db.Create(&expense1)
	db.Create(&expense2)
}
