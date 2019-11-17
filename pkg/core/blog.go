package core

// Blog defines the API interface
type Blog interface {
	FindPost(id string) (*Post, error)
	FindAllPosts() ([]*Post, error)
	DeletePost(post *Post) error
	UpdatePost(post *Post) error
	CreatePost(post ...*Post) error
	CountPosts() (int, error)
}
