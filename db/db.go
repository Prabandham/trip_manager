package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	db *gorm.DB
}

func Connection() *gorm.DB {
	cdb, err := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db := &DB{db: cdb}
	return db.db
}
