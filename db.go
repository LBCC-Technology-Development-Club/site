package db

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
func AllPosts() []post.Post {
	db := connect()
	selDB, err := db.Query("SELECT * FROM Post")
	if err != nil {
		panic(err.Error())
	}
	post := post.Post{}
	res := []post.Post{}
	for selDB.Next() {
		var uID, pID int
		var title, summary, body, timestamp string
		err = selDB.Scan(&pID, &timestamp, &title, &summary, &body, &uID)
		if err != nil {
			panic(err.Error())
		}
		post.UserID = uID
		post.PostID = pID
		post.Title = title
		post.Summary = summary
		post.Body = body
		post.Timestamp = timestamp
		res = append(res, post)
	}

	return res
}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/post/{pID}", GetPost)
	router.Get("/post/", GetAllPosts)
	router.Delete("/post/{pID}", DeletePost)
	router.Post("/post/", CreatePost)

	router.Post("/login/", Login)

	return router
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")

	selDB, err := db.Query("SELECT * FROM Post WHERE pID = " + pID); err != nil {
		panic(err.Error())
	}

	post := post.Post{}

	for selDB.Next() {
		var uID, pID int
		var title, summary, body, timestamp string
		err = selDB.Scan(&pID, &timestamp, &title, &summary, &body, &uID)
		if err != nil {
			panic(err.Error())
		}
		post.UserID = uID
		post.PostID = pID
		post.Title = title
		post.Summary = summary
		post.Body = body
		post.Timestamp = timestamp
	}

	render.JSON(w, r, post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := connect()

	selDB, err := db.Query("SELECT * FROM Post"); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	post := post.Post{}
	res := []post.Post{}

	for selDB.Next() {
		var uID, pID int
		var title, summary, body, timestamp string
		err = selDB.Scan(&pID, &timestamp, &title, &summary, &body, &uID)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}
		post.UserID = uID
		post.PostID = pID
		post.Title = title
		post.Summary = summary
		post.Body = body
		post.Timestamp = timestamp
		res = append(res, post)
	}

	render.JSON(w, r, res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := connect()

	pID := chi.URLParam(r, "pID")
	response := make(map[string]string)

	selDB, err := db.Query("DELETE FROM Post WHERE pID = " + pID); err != {
		response["message"] = "Failed to delete post"
		log.Panicf("Logging error: %s\n", err.Error())
	} else {
		response["message"] = "Deleted post successfully"
	}
	
	render.JSON(w, r, response)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	db := connect()

	response := make(map[string]string)

	selDB, err := db.Query("INSERT")
}