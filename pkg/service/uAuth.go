package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

var hmacSecret []byte
var initVar bool = false

// uAuthInit reads the secret from the server
func uAuthInit() {
	hmacSecret = []byte(os.Getenv("HMACSECRET"))
	initVar = true
}

// LoginRoutes defines the endpoints for user login and auth
func LoginRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{email}/{password}", Login)
	router.Delete("/", Logout)
	router.Post("/signup", Signup)

	return router
}

// Login tries to log a user in, and provide them with a JWT
func Login(w http.ResponseWriter, r *http.Request) {
	if !initVar {
		uAuthInit()
	}
	// Check db for a user with the correct email and password
	db := Connect()

	email := chi.URLParam(r, "email")
	password := chi.URLParam(r, "password")

	query := `SELECT User.saltedhash, User.uID, User.isAdmin FROM User WHERE User.email = "` + email + `";`

	selDB, err := db.Query(query)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	var hash string
	var uID, admin int
	for selDB.Next() {
		err = selDB.Scan(&hash, &uID)
		if err != nil {
			log.Panicf("Logging error: %s\n", err.Error())
		}
	}

	response := make(map[string]string)

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf("Login error: %s\n", err.Error())
		response["message"] = "Invalid login"
		render.JSON(w, r, response)
		db.Close()
		return
	}

	response["message"] = "Logged in"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": uID,
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	response["jwt"] = tokenString

	if admin == 0 {
		response["admin"] = "false"
	}
	if admin == 1 {
		response["admin"] = "true"
	}
	render.JSON(w, r, response)
	db.Close()
	// Return a JWT for continued authentication
}

// Logout makes the JWT for a user non-valid
func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete the JWT that the user is using
}

// Signup makes a new user, and logs them in with a JWT
func Signup(w http.ResponseWriter, r *http.Request) {
	if !initVar {
		uAuthInit()
	}
	// Check that there are no users that have the same email
	db := Connect()

	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	query := `SELECT COUNT(*) From User WHERE User.email = "` + user.Email + `";`

	selDB, err := db.Query(query)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	response := make(map[string]string)

	for selDB.Next() {
		var num int
		err = selDB.Scan(&num)

		if num != 0 {
			response["message"] = "A user with this email already exists"
			render.JSON(w, r, response)
			db.Close()
			return
		}
	}

	// Salt and hash the password and store the user
	password := []byte(user.Password)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	user.Hash = string(hash)
	query = `INSERT INTO User(isAdmin, uID, Name, email, saltedhash, role) VALUES (0, NULL, "` + user.Name + `", "` + user.Email + `", "` + user.Hash + `", "Member");`

	selDB, err = db.Query(query)
	if err != nil {
		response["message"] = "Error signing up"
		log.Panicf("Logging error: %s\n", err.Error())
	} else {
		response["message"] = "Successfuly signed up"
	}

	// Provide the user with a JWT to log them in
	render.JSON(w, r, response)
	db.Close()
}
