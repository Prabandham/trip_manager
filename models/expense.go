package models

type Expense struct {
	ID          int    `db:id`
	TripId      int    `db:"trip_id"`
	PersonId    int    `db:"person_id"`
	Label       string `db:"label"`
	Description string `db:"description"`
	Spent       int    `db:"spent"`
	CreatedAt   string `db:"created_at"`
}

var ExpenseSchema string = `
	CREATE TABLE expenses (
		id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
		trip_id int,
		person_id int,
		label varchar(255),
		description text,
		spent int,
		created_at varchar(255),
		FOREIGN KEY (trip_id) REFERENCES trips(id),
		FOREIGN KEY (person_id) REFERENCES people(id)
	) ENGINE=InnoDB;
`

var ExpenseIndexes = []string{"ALTER TABLE expenses ADD INDEX expenses_trip_id (trip_id);", "ALTER TABLE expenses ADD INDEX expenses_person_id (person_id);"}
