package models

import (
	"errors"
	"github.com/Prabandham/trip_manager/db"
)

const (
	PersonSchema string = `
	CREATE TABLE people (
		id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
		user_name varchar(255),
		phone_number varchar(10),
		confirmation_code varchar(5),
		confirmed varchar(3),
		created_at varchar(255)
	) ENGINE=InnoDB;
`
)

var PersonIndexes = []string{`ALTER TABLE people ADD UNIQUE INDEX people_unique_phone (phone_number);`, `ALTER TABLE people ADD INDEX people_confirmed_phone (confirmed, phone_number);`}

// Person - we will save created_at only when the person
// confirms his code.
type Person struct {
	Id               int    `db:"id"`
	UserName         string `db:"user_name" json:"user_name"`
	PhoneNumber      string `db:"phone_number" json:"phone_number"`
	ConfirmationCode string `db:"confirmation_code" json:"-"`
	Confirmed        string `db:"confirmed" json:"confirmed"`
	CreatedAt        string `db:"created_at" json:"-"`
}

func (p *Person) Validate() (*Person, error) {
	if p.PhoneNumber == "" {
		return nil, errors.New("Phone Number not entered")
	}
	var dbUser Person
	var query_error error

	if p.ConfirmationCode == "" {
		query_error = db.Connection.Get(&dbUser, "SELECT * FROM people where people.phone_number=?", p.PhoneNumber)
	} else {
		query_error = db.Connection.Get(&dbUser, "SELECT * FROM people where people.phone_number=? and people.confirmation_code=?", p.PhoneNumber, p.ConfirmationCode)
	}
	return &dbUser, query_error
}
