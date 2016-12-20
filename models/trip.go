package models

type Trip struct {
	Id          int    `db:"id"`
	Destination string `db:"destination"`
	Description string `db:"Description"`
	StartDate   string `db:"start_date"`
	EndDate     string `db:end_date`
	CreatedAt   string `db:created_at`
}

type TripPeople struct {
	TripId   int `db:"trip_id"`
	PersonId int `db:"person_id"`
}

const (
	TripSchema string = `
	CREATE TABLE trips (
		id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
		destination varchar(255),
		description text,
		start_date varchar(255),
		end_date varchar(255),
		created_at varchar(255)
	) ENGINE=InnoDB;
	`
	TripPeopleSchema string = `
	CREATE TABLE trip_people (
		trip_id int,
		person_id int,
		FOREIGN KEY (trip_id) REFERENCES trips(id),
		FOREIGN KEY (person_id) REFERENCES person(id)
	) ENGINE=InnoDB;
`
)

var TripIndexes = []string{"ALTER TABLE trips ADD INDEX trips_destination (destination);", "ALTER TABLE trips ADD INDEX trips_start_date  (start_date);"}
var TripPeopleIndexes = []string{"ALTER TABLE trip_people ADD INDEX tp_trip_id (trip_id);", "ALTER TABLE trip_people ADD INDEX tp_person_id (person_id);"}
