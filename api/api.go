package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Post struct {
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
}

// BlogRoutes defines the endpoints for /blog
func BlogRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{postID}", GetAPost)
	router.Delete("/{postID}", DeletePost)
	router.Post("/", CreatePost)
	router.Get("/", GetAllPosts)
	return router
}

// GetAPost renders a post from the postID in the URL
func GetAPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	post := Post{
		Slug:    postID,
		Title:   "First Post",
		Summary: "This is a summary",
		Body:    "This is the body",
	}
	render.JSON(w, r, post)
}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Post successfully"
	render.JSON(w, r, response)
}

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Post successfully"
	render.JSON(w, r, response)
}

// GetAllPosts returns all the posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{
			Slug:    "slug",
			Title:   "hello world",
			Summary: "hello summary",
			Body:    "hello world from go",
		},
	}
	render.JSON(w, r, posts)
}

// Routes builds the base routes for the API
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/", func(r chi.Router) {
		r.Mount("/api/blog", BlogRoutes())
	})

	return router
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
