package service

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// BlogRoutes defines the API endpoints for the /blog
func BlogRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/post/{pID}", GetPost)
	router.Get("/post", GetAllPosts)
	router.Delete("/post/{pID}", DeletePost)
	router.Post("/post", CreatePost)
	//router.Put("/post/{pID}", UpdatePost)

	router.Get("/comment/{pID}", GetPostComments)
	//router.Post("/comment/{pID}", CreateComment)
	//router.Delete("/comment/{cID}", DeleteComment)

	return router
}

// GetPost gets a specific post from the database
func GetPost(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")

	selDB, err := db.Query("SELECT * FROM Post WHERE pID = " + pID)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	post := Post{}

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
	}

	render.JSON(w, r, post)
}

// GetAllPosts gets all the posts from the database
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	selDB, err := db.Query("SELECT * FROM Post")
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	post := Post{}
	res := []Post{}

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

// GetPostComments gets all the comments for a post
func GetPostComments(w http.ResponseWriter, r *http.Request) {

}

/*
Below this line requires user authentication
*/

// DeletePost deletes a specific post from the database
func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")
	response := make(map[string]string)

	_, err := db.Query("DELETE FROM Post WHERE pID = " + pID)
	if err != nil {
		response["message"] = "Failed to delete post"
		log.Panicf("Logging error: %s\n", err.Error())
	} else {
		response["message"] = "Deleted post successfully"
	}

	render.JSON(w, r, response)
}

// CreatePost makes a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	//db := Connect()

	//response := make(map[string]string)

	//selDB, err := db.Query("INSERT")
}
