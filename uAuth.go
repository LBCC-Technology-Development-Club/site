package uAuth

import (
	"github.com/LBCC-Technology-Development-Club/site/database"
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", Login)
	router.Delete("/", Logout)
	router.Post("/signup/", Signup)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	// Check db for a user with the correct email and password

	// Return a JWT for continued authentication
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete the JWT that the user is using
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// Check that there are no users that have the same email

	// Salt and hash the password and store the user

	// Provide the user with a JWT to log them in
}