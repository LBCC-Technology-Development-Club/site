package post

// Post represents a blog post
type Post struct {
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Body      string `json:"body"`
	Timestamp string `json:"timestamp"`
}
