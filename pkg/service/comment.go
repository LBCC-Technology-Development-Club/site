package service

// Comment represents a comment on a post
type Comment struct {
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	CommentID int    `json:"comment_id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
