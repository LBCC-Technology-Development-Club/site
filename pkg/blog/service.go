package blog

import (
	"strconv"

	"github.com/LBCC-Technology-Development-Club/site/pkg/core"
)

// RawPost describes the structure of raw posts that we get from the client
type RawPost struct {
	PostID  int64  `json:"pid"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
	// Add timestamp here
}

// Service stores current user ID and blog
type Service struct {
	userID string
	blog   core.Blog
}

// GetPosts returns all posts
func (s Service) GetPosts() ([]*core.Post, error) {
	return s.blog.FindAllPosts()
}

// CreatePost creates a new post
func (s Service) CreatePost(rawPost RawPost) (*core.Post, error) {
	post := s.rawPostToPost(rawPost)
	err := s.blog.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
