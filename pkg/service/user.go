package service

// User defines a user's data
type User struct {
	uID        int    `json: user_id`
	Name       string `json: name`
	Email      string `json: email`
	IsAdmin    bool   `json: is_admin`
	saltedHash string `json: hash`
	role       string `json: role`
}
