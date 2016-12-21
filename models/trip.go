package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Trip struct {
	gorm.Model
	Destination string     `db:"destination" json:"destination"`
	Description string     `db:"description" json:"description"`
	StartDate   time.Time  `db:"start_date" json:"start_date"`
	EndDate     *time.Time `db:"end_date" json:"end_date"`

	People   []Person `gorm:"many2many:trip_people"`
	Moments  []Moment
	Expenses []Expense
}

type TripPeople struct {
	gorm.Model
	TripId   int `db:"trip_id"`
	PersonId int `db:"person_id"`
}

func (TripPeople) TableName() string {
	return "trip_people"

}
