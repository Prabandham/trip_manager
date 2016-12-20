package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DbConnection struct{}

var Connection *sqlx.DB

func init() {
	if Connection == nil {
		Connection = sqlx.MustConnect("mysql", "root:root@/trip_manager?parseTime=true")
		// TODO check what is the best and then use those numbers
		Connection.SetMaxIdleConns(50)
		Connection.SetMaxOpenConns(5000)
	}
}
