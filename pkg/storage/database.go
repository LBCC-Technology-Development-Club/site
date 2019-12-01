package storage

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

func connect() *sql.DB {
	config := mysql.NewConfig()

	config.User = os.Getenv("DBUSER")
	config.Passwd = os.Getenv("DBPASSWORD")
	config.Net = "tcp"
	config.Addr = "localhost:3306"
	config.DBName = "cs340"

	dsn := config.FormatDSN()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err.Error())
	}

	return db
}
