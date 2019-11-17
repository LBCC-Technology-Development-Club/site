package core

// Post represents a blog post
type Post struct {
	UserID  string `json:"user_id"`
	PostID  string `json:"post_id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
	// Add timestamp here
}
