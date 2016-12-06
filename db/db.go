package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
}

var Connection gorm.DB

//This sets up a global variable with a Connection
func Initialize() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10000)

	Connection = *db
	return &Connection
}
