package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Moment struct {
	gorm.Model
	FileType     int    `db:"file_type"`
	FileSize     int    `db:"file_size"`
	FileLocation string `db:"file_location"`
	Location     string `db:"location"`
	Trip         Trip
	TripId       int
}
