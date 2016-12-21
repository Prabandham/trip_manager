// This is going to be a one time call.
// To setup database and tables and indexes.
package main

import (
	. "github.com/Prabandham/trip_manager/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func LoadSchema() {
	db, _ := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Person{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Trip{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Expense{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Moment{})

	db.Model(&Person{}).AddUniqueIndex("idx_person_phone_number", "phone_number")
	db.Model(&Trip{}).AddIndex("idx_trip_destination", "destination")
	db.Model(&Expense{}).AddIndex("idx_expense_person_id", "person_id")
	db.Model(&Expense{}).AddIndex("idx_expense_trip_id", "trip_id")
	db.Model(&Expense{}).AddIndex("idx_expense_trip_id_person_id", "trip_id", "person_id")
	db.Model(&Moment{}).AddIndex("idx_moment_trip_id", "trip_id")
	db.Model(&Moment{}).AddIndex("idx_moment_file_type", "file_type")
}
