package models

type Moment struct {
	Id           int    `db:"id"`
	TripId       int    `db:"trip_id"`
	FileType     int    `db:"file_type"`
	FileSize     int    `db:"file_size"`
	FileLocation string `db:"file_location"`
	Location     string `db:"location"`
	CreatedAt    string `db:"created_at"`
}

var MomentsSchema string = `
	CREATE TABLE moments (
		id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
		trip_id int,
		file_type varchar(255),
		file_size int,
		file_location varchar(255),
		location varchar(255),
		created_at varchar(255),
		FOREIGN KEY (trip_id) REFERENCES trips(id)
	) ENGINE=InnoDB;
`

var MomentsIndexes = []string{"ALTER TABLE expenses ADD INDEX expenses_spent (spent);", "ALTER TABLE expenses ADD INDEX expenses_person_spent (person_id, spent);", "ALTER TABLE moments ADD INDEX moments_trip_id (trip_id);"}
