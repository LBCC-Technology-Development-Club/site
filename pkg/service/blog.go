package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

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
	router.Get("/post/{pID}/tags", GetPostTags)
	router.Put("/post/{pID}", SetPostTags)

	router.Get("/comment/{pID}", GetPostComments)
	router.Post("/comment/{pID}", CreateComment)
	router.Delete("/comment/{cID}", DeleteComment)

	router.Get("/user/{uID}", GetUser)
	router.Get("/user/{uID}/posts", GetUserPosts)

	return router
}

// GetPost gets a specific post from the database
func GetPost(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")

	query := "SELECT Post.*, User.Name FROM Post, User WHERE Post.pID = " + pID + " AND Post.uID = User.uID"

	selDB, err := db.Query(query)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	post := Post{}

	for selDB.Next() {
		var uID, pID int
		var title, summary, body, timestamp, author string
		err = selDB.Scan(&pID, &timestamp, &title, &summary, &body, &uID, &author)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}
		post.UserID = uID
		post.PostID = pID
		post.Author = author
		post.Title = title
		post.Summary = summary
		post.Body = body
		post.Timestamp = timestamp
	}

	query = `SELECT Post_Tags.tag FROM Post_Tags WHERE Post_Tags.pID = "` + strconv.Itoa(post.PostID) + `";`
	tags := []string{}

	selDB, err = db.Query(query)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	for selDB.Next() {
		var tag string
		err = selDB.Scan(&tag)
		tags = append(tags, tag)
	}

	post.Tags = tags

	db.Close()
	render.JSON(w, r, post)
}

// GetAllPosts gets all the posts from the database
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	selDB, err := db.Query("SELECT Post.*, User.Name FROM Post, User WHERE Post.uID = User.uID ORDER BY Post.pID DESC")
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	post := Post{}
	res := []Post{}

	for selDB.Next() {
		var uID, pID int
		var title, summary, body, timestamp, author string
		err = selDB.Scan(&pID, &timestamp, &title, &summary, &body, &uID, &author)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}
		post.UserID = uID
		post.PostID = pID
		post.Author = author
		post.Title = title
		post.Summary = summary
		post.Body = body
		post.Timestamp = timestamp

		res = append(res, post)
	}
	db.Close()
	render.JSON(w, r, res)
}

// GetPostTags gets all the tags associated with each post
func GetPostTags(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")

	selDB, err := db.Query(`SELECT Post_Tags.tag FROM Post_Tags WHERE Post_Tags.pID = "` + pID + `";`)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	res := []string{}

	for selDB.Next() {
		var tag string
		err = selDB.Scan(&tag)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}
		res = append(res, tag)
	}

	db.Close()
	render.JSON(w, r, res)
}

// GetPostComments gets all the comments for a post
func GetPostComments(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")

	query := "SELECT Comment.*, User.Name FROM Comment, User WHERE Comment.pID = " + pID + " AND Comment.uID = User.uID"

	selDB, err := db.Query(query)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	comment := Comment{}
	res := []Comment{}

	for selDB.Next() {
		var uID, pID, cID int
		var timestamp, content, author string
		err = selDB.Scan(&cID, &timestamp, &content, &uID, &pID, &author)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}

		comment.CommentID = cID
		comment.Timestamp = timestamp
		comment.Content = content
		comment.UserID = uID
		comment.PostID = pID
		comment.Author = author

		res = append(res, comment)
	}

	db.Close()
	render.JSON(w, r, res)
}

// GetUser returns the user data
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	uID := chi.URLParam(r, "uID")

	user := User{}

	query := "SELECT isAdmin, Name, role FROM User WHERE uID = " + uID

	selDB, err := db.Query(query)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	for selDB.Next() {
		var isAdmin bool
		var Name, role string
		err = selDB.Scan(&isAdmin, &Name, &role)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}

		user.IsAdmin = isAdmin
		user.Name = Name
		user.Role = role
	}

	db.Close()
	render.JSON(w, r, user)
}

// GetUserPosts finds the posts for a specific user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	uID := chi.URLParam(r, "uID")

	selDB, err := db.Query(`SELECT DISTINCT Post.*, User.Name FROM Post, User WHERE Post.uID = User.uID AND Post.uID = "` + uID + `" ORDER BY Post.pID DESC;`)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	post := Post{}
	res := []Post{}

	for selDB.Next() {
		var uID, pID int
		var title, summary, body, timestamp, author string
		err = selDB.Scan(&pID, &timestamp, &title, &summary, &body, &uID, &author)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}
		post.UserID = uID
		post.PostID = pID
		post.Author = author
		post.Title = title
		post.Summary = summary
		post.Body = body
		post.Timestamp = timestamp

		res = append(res, post)
	}

	db.Close()
	render.JSON(w, r, res)
}

/*
Below this line requires user authentication
*/

// DeletePost deletes a specific post from the database
func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")

	// Check that the post is not already deleted

	commentQuery := "DELETE FROM Comment WHERE Comment.pID = " + pID + ";"
	tagsQuery := "DELETE FROM Post_Tags WHERE Post_Tags.pID = " + pID + ";"
	stickyQuery := "DELETE FROM Stickied_Posts WHERE Stickied_Posts.pID = " + pID + ";"
	query := "DELETE FROM Post WHERE Post.pID = " + pID + ";"

	response := make(map[string]string)

	_, err := db.Query(commentQuery)
	if err != nil {
		response["message"] = "Could not delete comments for this post"
		log.Panicf("Logging error: %s\n", err.Error())
	}

	_, err = db.Query(tagsQuery)
	if err != nil {
		response["message"] = "Could not delete tags for this post"
		log.Panicf("Logging error: %s\n", err.Error())
	}

	_, err = db.Query(stickyQuery)
	if err != nil {
		response["message"] = "Could not delete sticky post tag for this post"
		log.Panicf("Logging error: %s\n", err.Error())
	}

	_, err = db.Query(query)
	if err != nil {
		response["message"] = "Could not delete this post"
		log.Panicf("Logging error: %s\n", err.Error())
	}

	render.JSON(w, r, response)
	db.Close()
}

// DeleteComment deletes a specific post from the database
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	cID := chi.URLParam(r, "cID")
	response := make(map[string]string)

	_, err := db.Query("DELETE FROM Comment WHERE cID = " + cID)
	if err != nil {
		response["message"] = "Failed to delete comment"
		log.Panicf("Logging error: %s\n", err.Error())
	} else {
		response["message"] = "Deleted comment successfully"
	}

	db.Close()
	render.JSON(w, r, response)
}

// CreatePost makes a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	post := Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	// CHANGE THIS TO BE CLIENT SIDE
	post.Timestamp = time.Now().String()

	query := `INSERT INTO Post(pID, timestamp, title, summary, body, uID) VALUES (NULL, "` + string(post.Timestamp) + `", "` + post.Title + `", "` + post.Summary + `", "` + post.Body + `", "` + string(post.UserID) + `")`

	response := make(map[string]string)

	_, err = db.Query(query)
	if err != nil {
		response["message"] = "Failed to create post"
		log.Panicf("Logging error: %s\n", err.Error())
	} else {
		response["message"] = "Created post successfully"
	}

	db.Close()
	render.JSON(w, r, response)
}

// SetPostTags sets the tags on a specific post
func SetPostTags(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	// verify the user

	tags := Tags{}

	pID := chi.URLParam(r, "pID")

	err := json.NewDecoder(r.Body).Decode(&tags)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}
	response := make(map[string]string)

	newTags := SliceDiff(tags.NewTags, tags.OldTags)
	deletedTags := SliceDiff(tags.OldTags, tags.NewTags)

	for _, tag := range newTags {
		query := `INSERT INTO Post_Tags(pID, tag) VALUES ("` + string(pID) + `", "` + tag + `");`

		_, err = db.Query(query)
		if err != nil {
			response["message"] = "Failed to add tag"
			log.Panicf("Logging error: %s\n", err.Error())
		} else {
			response["message"] = "Added tag successfully"
		}
	}

	for _, tag := range deletedTags {
		query := `DELETE FROM Post_Tags WHERE Post_Tags.pID = "` + string(pID) + `" AND Post_Tags.tag = "` + tag + `";`

		_, err = db.Query(query)
		if err != nil {
			response["message"] = "Failed to add tag"
			log.Panicf("Logging error: %s\n", err.Error())
		} else {
			response["message"] = "Added tag successfully"
		}
	}

	db.Close()
	render.JSON(w, r, response)
}

// CreateComment makes a new post
func CreateComment(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	pID := chi.URLParam(r, "pID")

	comment := Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	comment.Timestamp = time.Now().String()

	query := `INSERT INTO Comment(cID, timestamp, content, uID, pID) VALUES (NULL, "` + string(comment.Timestamp) + `", "` + comment.Content + `", "` + string(comment.UserID) + `", "` + pID + `")`
	log.Printf("%s\n", query)
	response := make(map[string]string)

	_, err = db.Query(query)
	if err != nil {
		response["message"] = "Failed to create comment"
		log.Panicf("Logging error: %s\n", err.Error())
	} else {
		response["message"] = "Created comment successfully"
	}

	db.Close()
	render.JSON(w, r, response)
}

/* Helpers */

// Find returns if an object exists in a slice, and the index of the object
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// SliceDiff returns the elements in a that are not in b
func SliceDiff(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
