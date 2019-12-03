package blog

import (
	"github.com/go-chi/chi"
	"github.com/LBCC-Technology-Development-Club/site/database"
)


func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/post/{pID}", GetPost)
	router.Get("/post/", GetAllPosts)
	router.Delete("/post/{pID}", DeletePost)
	router.Post("/post/", CreatePost)
	router.Update("/post/{pID}", UpdatePost)

	router.Get("/comment/{pID}", GetPostComments)
	router.Post("/comment/{pID}", CreateComment)
	router.Delete("/comment/{cID}", DeleteComment)

	return router
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

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
	db := database.Connect()

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
	db := database.Connect()

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
	db := database.Connect()

	response := make(map[string]string)

	selDB, err := db.Query("INSERT")
}