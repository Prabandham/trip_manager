package models

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Person - we will save created_at only when the person
// confirms his code.
type Person struct {
	gorm.Model
	UserName         string `db:"user_name" json:"user_name"`
	PhoneNumber      string `db:"phone_number" json:"phone_number"`
	ConfirmationCode string `db:"confirmation_code" json:"confirmation_code"`
	Confirmed        bool   `db:"confirmed" json:"confirmed"`

	Trips []Trip `gorm:"many2many:trip_people"`
}

func (p *Person) Validate() (*Person, error) {
	db, _ := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close()

	if p.PhoneNumber == "" {
		return new(Person), errors.New("Phone Number not entered")
	}
	var dbUser Person

	if p.ConfirmationCode == "" {
		db.Where("phone_number = ?", p.PhoneNumber).First(&dbUser)
		return &dbUser, nil
	} else {
		db.Where("phone_number = ? AND confirmation_code = ?", p.PhoneNumber, p.ConfirmationCode).First(&dbUser)
		if dbUser.ConfirmationCode == p.ConfirmationCode {
			return &dbUser, nil
		} else {
			return new(Person), errors.New("Confirmation Code does not match")
		}
	}
	return new(Person), nil
}

func (p *Person) Save() (*Person, error) {
	db, _ := gorm.Open("mysql", "root:root@/trip_manager?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close()

	if db.NewRecord(p) {
		db.Create(&p)
		return p, nil
	}

	db.Model(&p).Update(&p)
	return p, nil
}
