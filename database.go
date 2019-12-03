package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LBCC-Technology-Development-Club/site/post"
	"github.com/go-chi/chi"
	"github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	config := mysql.NewConfig()

	// Set up the DSN with environment vars
	config.User = os.Getenv("DBUSER")
	config.Passwd = os.Getenv("DBPASSWORD")
	config.Net = "tcp"
	config.Addr = os.Getenv("DBADDR")
	config.DBName = "cs340"

	dsn := config.FormatDSN()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	
	return db
}