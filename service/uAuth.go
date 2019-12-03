package service

import (
	"net/http"

	"github.com/go-chi/chi"
)

// LoginRoutes defines the endpoints for user login and auth
func LoginRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", Login)
	router.Delete("/", Logout)
	router.Post("/signup", Signup)

	return router
}

// Login tries to log a user in, and provide them with a JWT
func Login(w http.ResponseWriter, r *http.Request) {
	// Check db for a user with the correct email and password

	// Return a JWT for continued authentication
}

// Logout makes the JWT for a user non-valid
func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete the JWT that the user is using
}

// Signup makes a new user, and logs them in with a JWT
func Signup(w http.ResponseWriter, r *http.Request) {
	// Check that there are no users that have the same email

	// Salt and hash the password and store the user

	// Provide the user with a JWT to log them in
}
