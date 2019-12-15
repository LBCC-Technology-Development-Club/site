package service

// User defines a user's data
type User struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	Hash     string `json:"hash"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
