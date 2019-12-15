package service

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Connect opens a connection to the database
func Connect() *sql.DB {
	config := mysql.NewConfig()

	// Set up the DSN with environment vars
	config.User = os.Getenv("DBUSER")
	config.Passwd = os.Getenv("DBPASSWORD")
	config.Net = "tcp"
	config.Addr = os.Getenv("DBADDR")
	config.DBName = os.Getenv("DBNAME")

	dsn := config.FormatDSN()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	return db
}
