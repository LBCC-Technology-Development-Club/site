package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LBCC-Technology-Development-Club/site/pkg/core"
	"github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	config := mysql.NewConfig()

	config.User = os.Getenv("DBUSER")
	config.Passwd = os.Getenv("DBPASSWORD")
	config.Net = "tcp"
	config.Addr = os.Getenv("DBADDR")
	config.DBName = "cs340"

	dsn := config.FormatDSN()

	fmt.Printf("%s\n", dsn)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// AllPosts gets all posts from db
func AllPosts() {
	db := connect()

	selDB, err := db.Query("SELECT * FROM Post")
	if err != nil {
		panic(err.Error())
	}
	post := core.Post{}
	res := []core.Post{}
	for selDB.Next() {
		var uID, pID int
		var title, summary, body, timestamp string
		err = selDB.Scan(&pID, &timestamp, &title, &summary, &body, &uID)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("uID: %d\n", uID)
		post.UserID = uID
		fmt.Printf("pID: %d\n", pID)
		post.PostID = pID
		fmt.Printf("Title: %s\n", title)
		post.Title = title
		post.Summary = summary
		post.Body = body
		post.Timestamp = timestamp
		res = append(res, post)
	}
}

func main() {
	AllPosts()
}
